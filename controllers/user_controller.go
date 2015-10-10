package controllers

import (
	"strings"
	"beego_study/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Users() {
	userId, _ := c.GetInt(":id");
	var user, err = models.GetUser(userId);
	if err != nil {
		c.Data["json"] = nil;
	}else {
		c.Data["json"] = user;
	}
	c.ServeJson()
	strings.Split("", ":")

}