package controllers

import (
	"beego_study/models"
	"strings"
	"beego_study/entities"
	"log"
	"github.com/astaxie/beego"
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
	log.Println("QQ回调登录")
	nickname := c.GetString("nickname")
	qcOpenId := c.GetString("qcOpenId")
	sex := c.GetString("sex")
	user,_ := models.FundQCUser(qcOpenId)
	c.Data["openId"] = qcOpenId
	//用户存在,密码为空,跳转到输入密码的页面,这一步用来保护用户账号安全的,暂时先不添加了.
	if nil != user && strings.EqualFold(qcOpenId,user.QcOpenId) {
		//跳转到设置密码的页面
		c.SetSession("user", user)
		c.Data["user"] = user
		c.Ctx.Redirect(302,"/")

	}else{
		log.Println("QQ回调登录,新创建用户对象")
		var err error
		 user,err:= models.CreateQCUser(nickname+"_"+qcOpenId,qcOpenId,sex)
		if err == nil {
			log.Printf("创建的新用户name:%s",user.Name)
			c.SetSession("cacheUser",user);

			//跳转到设置密码的页面
			c.Data["showRightBar"] = false
			c.TplNames = "pwd.html"
		}else{
			c.TplNames = "login.html"
		}
	}
}

func (c *UserController) IgnorePwdLogin(){
	log.Println("用户提交密码信息")
	openId := c.GetString("opneId")
	pwd := c.GetString("pwd")
	repwd := c.GetString("repwd")
	u := c.GetSession("cacheUser")
	beego.Error("**************************u:",u)
	var user,ok = u.(*entities.User)
	if !ok {
		c.Data["showRightBar"] = false
		c.TplNames = "login.html"
	}
	log.Println("获取缓存的用户对象")
	log.Printf("参数信息openId:%s,user.OpenId:%s",openId,user.QcOpenId)
	log.Printf("用户的姓名信息name:%s",user.Name)
	//用户存在,并且请求匹配
	if len(openId) > 0 && strings.EqualFold(openId,user.QcOpenId)  {
		log.Println("缓存的用户open ID相等")
		if strings.EqualFold(pwd,repwd){
			log.Println("密码校验成功")
			user.Password = pwd
			models.SetUserPassword(user)
			c.SetSession("cacheUser",nil)
			c.SetSession("user",user)
			c.Data["user"] = user
			c.Ctx.Redirect(302,"/")
		}else{
			log.Println("密码校验不一致")
			c.Data["errMessage"] = "密码和重复密码不一致"
			c.TplNames = "pwd.html"
		}
	}else{
		log.Println("非法登录")
		//非法登录
		c.Ctx.Redirect(302,"/")
	}

}
