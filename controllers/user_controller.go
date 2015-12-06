package controllers

import (
	"beego_study/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Users() {
	userId, _ := c.GetInt(":id");
	var user, err = models.User(userId);
	if err != nil {
		c.Data["json"] = nil;
	}else {
		c.Data["json"] = user;
	}
	c.ServeJson()
}

func (c *UserController) Login() {
	c.Data["showRightBar"] = false
	c.TplNames = "login.html"
}

func (c *UserController) Session() {
	name := c.GetString("name")
	password := c.GetString("password")
	user, _ := models.FundUser(name, password)

	if (user.Valid(name,password)) {
		c.SetSession("user", user)
		c.Ctx.Redirect(302,"/")
	}else {
		c.SetError("用户名或者密码错误")
		c.TplNames = "login.html"
	}
}