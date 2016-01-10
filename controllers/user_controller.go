package controllers

import (
	"beego_study/models"
	"strings"
	"beego_study/entities"
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

//QC  回调页面
func (c*UserController) Redirect(){
	c.TplNames = "redirect.html"
}

func (c *UserController) QCSession(){
	nickname := c.GetString("nickname")
	qcOpenId := c.GetString("qcOpenId")
	sex := c.GetString("sex")
	user,_ := models.FundQCUser(qcOpenId)
	c.Data["openId"] = qcOpenId
	//用户存在,密码为空,跳转到输入密码的页面,这一步用来保护用户账号安全的,暂时先不添加了.
	if nil != user {
		//c.SetSession("cacheUser",user);
		//跳转到设置密码的页面
		c.SetSession("user", user)
		c.Data["user"] = user
		c.Ctx.Redirect(302,"/")

	}else{
		var err error
		 user,err:= models.CreateQCUser(nickname+"_"+qcOpenId,qcOpenId,sex)
		if err != nil {
			c.SetSession("cacheUser",user);
			//跳转到设置密码的页面
			c.Data["showRightBar"] = false
			c.TplNames = "pwd.html"
		}
	}
}

func (c *UserController) IgnorePwdLogin(){
	openId := c.GetString("opneId")
	pwd := c.GetString("pwd")
	repwd := c.GetString("repwd")
	user := c.GetSession("cacheuUser")
	if nil == user{
		c.Data["showRightBar"] = false
		c.TplNames = "login.html"
	}
	var u,ok = user.(entities.User)
	if !ok {
		c.Data["showRightBar"] = false
		c.TplNames = "login.html"
	}

	//用户存在,并且请求匹配
	if len(openId) > 0 && strings.EqualFold(openId,u.QcOpenId)  {
		if strings.EqualFold(pwd,repwd){
			u.Password = pwd
			models.SetUserPassword(&u)
			c.SetSession("cacheUser",nil)
			c.SetSession("user",u)
			c.Data["user"] = user
			c.Ctx.Redirect(302,"/")
		}else{
			c.Data["errMessage"] = "密码和重复密码不一致"
			c.TplNames = "pwd.html"
		}
	}else{
		//非法登录
		c.Ctx.Redirect(302,"/")
	}

}
