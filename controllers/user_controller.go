package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"github.com/astaxie/beego"
	"beego_study/exception"
	"beego_study/utils"
	"beego_study/utils/redis"
	"strings"
	"beego_study/constants"
	"fmt"
	"errors"
)

const
(
	COMMON_REGISTER = 0 //普通注册

	OAUTH_REGISTER = 1  //第三方授权注册
)

type UserController struct {
	BaseController
}

func (c *UserController) Users() {
	userId, _ := c.GetInt64(":id");
	var user, err = models.User(userId);
	if err != nil {
		c.Data["json"] = nil;
	}else {
		c.Data["json"] = user;
	}
	c.ServeJSON()
}

func (c *UserController) Login() {
	c.Data["showLeftBar"] = false
	c.TplName = "login.html"
}

func (c *UserController) OauthLogin() {
	c.Data["showLeftBar"] = false
	openUser := c.CurrentOpenUser()
	beego.Debug("openUser:", openUser)
	if nil == openUser || openUser.UserId != 0 {
		c.Ctx.Redirect(302, "/")
		return
	}
	c.TplName = "login_bind_open_user.html"
}

func (c *UserController) Logout() {
	c.DelSession("user")
	c.DelSession("openUser")
	c.Ctx.Redirect(302, "/")
}

func (c *UserController) Register() {
	c.Data["showLeftBar"] = false
	c.TplName = "register.html"
}

func (c *UserController) OauthRegister() {
	c.Data["showLeftBar"] = false
	openUser := c.CurrentOpenUser()
	beego.Debug("openUser:", openUser)
	if nil == openUser || openUser.UserId > 0 {
		c.Ctx.Redirect(302, "/")
		return
	}
	c.TplName = "register_bind_open_user.html"
}

func (c *UserController) CreateUser() {
	registerType, err := c.GetInt("registerType")
	if nil != err {
		c.Ctx.Redirect(302, "/")
	}

	if COMMON_REGISTER == registerType {
		c.commonCreateUser()
		return
	}

	if OAUTH_REGISTER == registerType {
		c.oauthCreateUser()
		return
	}
}

func (c *UserController) commonCreateUser() {

	var user entities.User
	user.Name = c.GetString("name")
	user.Mail = c.GetString("mail")
	user.Password = c.GetString("password")
	captcha := c.GetString("captcha")
	var err error
	isCaptchaValid := c.isMailCaptchaValid(captcha)
	if !isCaptchaValid {
		err = errors.New("验证码错误")
	}else {
		err = models.SaveUser(&user)
	}

	if err != nil {
		c.StringError(err.Error())
		c.Data["name"] = user.Name
		c.Data["mail"] = user.Mail
		c.Register()
	}

	c.SetCurrSession("user", user)

	c.Ctx.Redirect(302, "/")
}

func (c *UserController) oauthCreateUser() {

	var user entities.User
	user.Name = c.GetString("name")
	user.Mail = c.GetString("mail")
	user.Password = c.GetString("password")

	openUser := c.CurrentOpenUser()
	beego.Debug("openuser:", openUser)
	var unBindUser bool
	if nil != openUser && openUser.UserId == 0 {
		user = user.InitFromOpenUser(openUser)
		unBindUser = true
	}

	err := models.SaveUser(&user)

	if err == nil {
		c.SetCurrSession("user", user)
		c.Ctx.Redirect(302, "/")
		var err error
		var result int64
		if unBindUser {
			result, err = models.BindUserIdToOpenUser(openUser.OpenId, entities.OPEN_USER_TYPE_QQ, user.Id)
		}
		//防止单个openuser 绑定多个user
		if nil == err && result > 0 {
			openUser.UserId = user.Id
			c.SetSession("openUser", openUser)
		}
		return
	}

	c.StringError(err.Error())
	c.Data["name"] = user.Name
	c.Data["mail"] = user.Mail
	c.OauthRegister()
}

func (c *UserController) Session() {
	name := c.GetString("name")
	password := c.GetString("password")

	needCheckCaptcha := c.NeedCheckCaptcha()
	if needCheckCaptcha {
		result := c.VerifyCaptcha()
		if !result {
			c.RecordLoginFailTimes()
			c.StringError(exception.CAPTCHA_FALSE.Error())
			c.Data["showLeftBar"] = false
			c.TplName = "login.html"
			return
		}
	}

	user, err := models.FundUser(name, password)
	if err == nil {
		c.SetCurrSession("user", user)
		c.Ctx.Redirect(302, "/")
	}else {
		c.RecordLoginFailTimes()
		c.StringError(exception.USER_NAME_OR_PASS_UNMATCH.Error())
		c.Data["showLeftBar"] = false
		c.TplName = "login.html"
	}
}

func (c *UserController) OauthSession() {
	name := c.GetString("name")
	password := c.GetString("password")

	needCheckCaptcha := c.NeedCheckCaptcha()
	if needCheckCaptcha {
		result := c.VerifyCaptcha()
		if !result {
			c.RecordLoginFailTimes()
			c.StringError(exception.CAPTCHA_FALSE.Error())
			c.OauthLogin()
			return
		}
	}

	user, err := models.FundUser(name, password)
	if err == nil {
		c.SetCurrSession("user", user)
		c.Ctx.Redirect(302, "/")

		openUser := c.CurrentOpenUser()
		beego.Debug("openuser:", openUser)
		var result int64
		if nil != openUser && openUser.UserId == 0 {
			result, err = models.BindUserIdToOpenUser(openUser.OpenId, entities.OPEN_USER_TYPE_QQ, user.Id)
		}

		if nil == err && result > 0 {
			openUser.UserId = user.Id
			c.SetSession("openUser", openUser)
		}
		return
	}

	c.StringError(exception.USER_NAME_OR_PASS_UNMATCH.Error())

	c.RecordLoginFailTimes()

	c.OauthLogin()

}

func (c *UserController) CheckUserName() {
	name := c.GetString("name")
	err := models.CheckUserName(name)
	if nil != err {
		c.JsonError(err.Error())
	}else {
		c.JsonSuccess(nil)
	}

}

func (c *UserController) CheckUserMail() {
	mail := c.GetString("mail")
	err := models.CheckUserMail(mail)
	if nil != err {
		c.JsonError(err.Error())
	}else {
		c.JsonSuccess(nil)
	}

}

func (c *UserController) isMailCaptchaValid(captcha string) bool {
	sessionId := c.CruSession.SessionID()
	var captchaCache string
	mailCaptchaKey := strings.Replace(constants.MAIL_CAPTCHA_KEY, "$sessionId", sessionId, 1)
	redis_util.Get(mailCaptchaKey, &captchaCache)
	beego.Debug("captcha:", captcha, "captchaCache:", captchaCache)
	if (len(captcha) == 0 || len(captchaCache) == 0 || captchaCache != captcha) {
		return false
	}
	return true
}
func (c *UserController) CreateRegisterCaptcha() {
	mail := c.GetString("mail")
	sessionId := c.CruSession.SessionID()
	var captcha string
	mailCaptchaKey := strings.Replace(constants.MAIL_CAPTCHA_KEY, "$sessionId", sessionId, 1)
	redis_util.Get(mailCaptchaKey, &captcha)
	beego.Debug("captcha", captcha)
	if len(captcha) == 0 {
		captcha = utils.RandomIntCaptcha(6)
		redis_util.Set(mailCaptchaKey, captcha, 60)
	}
	content := fmt.Sprintf(models.MAIL_CAPTCHA_TEMPLATE, captcha, captcha)
	models.SendHtmlMail("threeperson(www.threeperson.com)注册验证码", content, []string{mail})
	c.JsonSuccess(60)

}

func (c *UserController) UserHome()  {
	pagination := c.NewPagination()
	userId := c.CurrentUserId()
	models.AllArticles(userId, pagination)
	models.SetLikeSign(pagination,userId)
	c.Data["pagination"] = pagination
	c.Data["user"]=c.CurrentUser()
	c.TplName = "user_home.html"
}

