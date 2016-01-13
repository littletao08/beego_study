package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	_"beego_study/initials"
)
func init() {
	beego.Router("", &controllers.IndexController{},"get:Index")

	beego.Router("/login", &controllers.UserController{},"get:Login")

	beego.Router("/users/session", &controllers.UserController{},"post:Session")


	//QQ 登录回调,这个页面只有js内容,会自动检测用户的登录信息并且异步把用户登录的信息发送到后台.
	beego.Router("/users/qq_redirect", &controllers.UserController{},"get:Redirect")

	//QQ用户登录,如果是老用户,直接登录成功
	beego.Router("/users/qclogin", &controllers.UserController{},"post:QCSession")

	//如果是新用户,要设置密码
	beego.Router("/users/dopwd", &controllers.UserController{},"post:IgnorePwdLogin")
	beego.Router("/articles", &controllers.ArticleController{},"get:Articles")
	beego.Router("/articles", &controllers.ArticleController{},"post:CreateArticle")
	beego.Router("/articles/new", &controllers.ArticleController{},"get:New")
	beego.Router("/articles/:id([0-9]", &controllers.ArticleController{},"get:ArticleDetail")
	beego.Router("/articles/:id([0-9]/edit", &controllers.ArticleController{},"get:EditArticle")
	beego.Router("/articles", &controllers.ArticleController{},"put:UpdateArticle")
	beego.Router("/categories/:category", &controllers.ArticleController{},"get:ArticlesGyCategory")
	beego.Router("/sponsors/new", &controllers.SponsorController{},"post:New")
	beego.Router("/articles/:id([0-9]/likes",&controllers.ArticleController{},"post:Like")

}


