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

func (c *UserController) Login()  {
	c.Data["showRightBar"] = false
	c.TplNames = "login.html"
}