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
	if (user.Name == name && user.Password == password) {
		c.SetSession("user", user)
		c.TplNames = "index.html"
	}else {
		response :=ResponseBody{Success:false,Message:"用户名或密码错误"}
		c.Data["showRightBar"] = false
		c.Data["response"] = response
		c.TplNames = "login.html"
	}
}