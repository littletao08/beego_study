package controllers

import (
	"github.com/astaxie/beego"
	"beego_study/utils"
	"beego_study/models"
	"errors"
	"beego_study/entities"
)

type OpenUserController struct {
	BaseController
}

func (c *OpenUserController) QqAuth() {
	params := make(map[string]string)
	params["client_id"] = models.AuthConfig.String("app_id")
	params["redirect_uri"] = models.AuthConfig.String("auth_redirect_uri")
	params["response_type"] = "code"
	var state = "111";
	params["state"] = state
	params["scope"] = models.AuthConfig.String("scope")

	var baseUrl = models.AuthConfig.String("authorize_url")

	subUrl := utils.BuildQueryString(params)

	var fullRequestStr = baseUrl + subUrl
	c.Redirect(fullRequestStr, 302)
}

func (c *OpenUserController) NewOrBindUser()  {
	c.Data["showLeftBar"] = false
	c.TplName="register.html"
}


func (c *OpenUserController) QqToken() {

	code := c.GetString("code")

	var loginPageUrl = "/users/login"
	//获取token
	tokenRes, err := models.QueryToken(code)
	beego.Debug("****************tokenRes:", tokenRes, "****************")
	if (nil != err ) {
		c.Redirect(loginPageUrl, 302)
		return
	}

	accessToken := tokenRes["access_token"]
	beego.Debug("****************accessToken:", accessToken, "****************")
	if len(accessToken) <= 0 {
		c.Redirect(loginPageUrl, 302)
		return
	}

	//获取openid
	openIdRes, err := models.QueryOpenId(accessToken)
	beego.Debug("****************openIdRes:", openIdRes, "****************")
	if (nil != err ) {
		beego.Error(err)
		c.Redirect("login.html", 302)
		return
	}

	openId := openIdRes["openid"]
	if len(openId) <= 0 {
		beego.Error(errors.New("openid["+openId+"] error"))
		c.Redirect(loginPageUrl, 302)
		return
	}

	//获取user_info
	openUser, err := models.OpenUserInfo(accessToken, openId)
	beego.Debug("err:",err,"userInfoRes:", openUser)

	if (nil != err ) {
		beego.Error(err)
		c.Redirect(loginPageUrl, 302)
		return
	}

	err = models.SaveOrUpdateOpenUser(openUser)
	if nil != err {
		beego.Error(err)
		c.Redirect(loginPageUrl, 302)
		return
	}

	openUser,_ =models.OpenUser(openId,entities.OPEN_USER_TYPE_QQ)
	//绑定了账号则跳转到首页,否则跳转到注册或绑卡页面
    if openUser.HasBindUser() {
		userId := openUser.UserId
		user,_ := models.User(userId)
		c.SetCurrSession("user", user)
	}else {
		c.SetSession("openUser",openUser)
		c.Redirect("/users/oauth_register", 302)
		return
	}

	c.Redirect("/", 302)
}



