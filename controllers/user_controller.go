package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"strings"
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
		c.Data["user"] = user
		c.Ctx.Redirect(302,"/")
	}else {
		c.StringError("用户名或者密码错误")
		c.Data["showRightBar"] = false
		c.TplNames = "login.html"
	}
}

func (c *UserController) QCSession(){
	nickname := c.GetString("nickname")
	qcOpenId := c.GetString("qcOpenId")
	sex := c.GetString("sex")
	user,_ := models.FundQCUser(qcOpenId)
	c.Data["openId"] = qcOpenId
	//用户存在,密码为空,跳转到输入密码的页面
	if len(user.Password) > 0  {
		//输入密码登录
		c.SetSession("cacheUser",user);
		c.Data["showRightBar"] = false
		c.TplNames = "pwd.html"
	} else if user != nil {
		c.SetSession("cacheUser",user);
		//跳转到设置密码的页面
	}else{
		var err error
		 user,err:= models.CreateQCUser(nickname+"_"+qcOpenId,qcOpenId,sex)
		if err != nil {
			c.SetSession("cacheUser",user);
			//跳转到设置密码的页面
		}
	}
}

func (c *UserController) IgnorePwdLogin(){
	openId := c.GetString("opneId")
	pwd := c.GetString("pwd")
	user := entities.User(c.GetSession("cacheUser"))
	//用户存在,并且请求匹配
	if len(openId) > 0 && strings.EqualFold(openId,user.QcOpenId)  {
		if c.GetBool("ignorePwd") {
			//忽略密码登录
			c.SetSession("user",user)
			//清空session
			c.SetSession("cacheUser",nil)
			c.Ctx.Redirect(302,"/")
		}else if strings.EqualFold(pwd,user.Password) {
			c.SetSession("user",user)
			//清空session
			c.SetSession("cacheUser",nil)
			c.Ctx.Redirect(302,"/")
		}else{
			//密码不正确
			c.Data["openId"]=openId;
		}
	}else{
		//非法登录
	}

}
