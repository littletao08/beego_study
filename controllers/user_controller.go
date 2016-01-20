package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"github.com/astaxie/beego"
	"beego_study/exception"
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

func (c *UserController) Logout() {
	c.DelSession("user")
	c.Ctx.Redirect(302, "/")
}

func (c *UserController) Register() {
	c.Data["showLeftBar"] = false
	c.TplName = "register.html"
}


func (c *UserController) CreateUser() {

	var user entities.User
	user.Name = c.GetString("name")
	user.Mail = c.GetString("mail")
	user.Password = c.GetString("password")

	openUser := c.CurrentOpenUser()
	beego.Debug("openuser:", openUser)
	var hasOpenUser bool
	if nil != openUser {
		user = user.InitFromOpenUser(openUser)
		hasOpenUser = true
	}

	err := models.SaveUser(&user)


	if err == nil{
		if hasOpenUser {
			models.BindUserIdToOpenUser(openUser.OpenId, entities.OPEN_USER_TYPE_QQ, user.Id)
		}
	} else {
		c.StringError(err.Error())
		c.Data["name"] = user.Name
		c.Data["mail"] = user.Mail
		c.Data["password"] = user.Password
		c.Register()
	}

	c.SetCurrSession("user",user)

	c.Ctx.Redirect(302, "/")
}

func (c *UserController) Session() {
	name := c.GetString("name")
	password := c.GetString("password")
	user, err := models.FundUser(name, password)

	if err == nil {
		c.SetCurrSession("user",user)
		c.Ctx.Redirect(302, "/")
	}else {
		c.StringError(exception.USER_NAME_OR_PASS_UNMATCH.Error())
		c.Data["showLeftBar"] = false
		c.TplName = "login.html"
	}
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