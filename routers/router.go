package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	_"beego_study/initials"
)
func init() {
	beego.Router("", &controllers.IndexController{},"get:Index")

	beego.Router("/users/login", &controllers.UserController{},"get:Login")
	beego.Router("/users/oauth_login", &controllers.UserController{},"get:OauthLogin")
	beego.Router("/users/logout", &controllers.UserController{},"get:Logout")
	beego.Router("/users/session", &controllers.UserController{},"post:Session")
	beego.Router("/users/oauth_session", &controllers.UserController{},"post:OauthSession")
	beego.Router("/users/register", &controllers.UserController{},"get:Register")
	beego.Router("/users/oauth_register", &controllers.UserController{},"get:OauthRegister")

	beego.Router("/users", &controllers.UserController{},"post:CreateUser")
	beego.Router("/users/check_user_name", &controllers.UserController{},"post:CheckUserName")
	beego.Router("/users/check_user_mail", &controllers.UserController{},"post:CheckUserMail")
	beego.Router("/users/check_user_mobile", &controllers.UserController{},"post:CheckUserMobile")


	beego.Router("/users/mob_register",&controllers.SmsController{},"post:Send")

	beego.Router("/articles", &controllers.ArticleController{},"get:Articles")
	beego.Router("/articles", &controllers.ArticleController{},"post:CreateArticle")
	beego.Router("/articles/new", &controllers.ArticleController{},"get:New")
	beego.Router("/articles/:id/edit", &controllers.ArticleController{},"get:EditArticle")
	beego.Router("/articles/:id([0-9]+)", &controllers.ArticleController{},"get:ArticleDetail")
	beego.Router("/articles", &controllers.ArticleController{},"put:UpdateArticle")
	beego.Router("/articles/:id([0-9]+)/likes",&controllers.ArticleController{},"post:Like")

	beego.Router("/categories/:category", &controllers.ArticleController{},"get:ArticlesGyCategory")

	beego.Router("/sponsors/new", &controllers.SponsorController{},"post:New")

	beego.Router("/open_users/:type/auth", &controllers.OpenUserController{},"get:QqAuth")
	beego.Router("/open_users/:type/token", &controllers.OpenUserController{},"get:QqToken")
	beego.Router("/open_users/mobile/mob_reg", &controllers.SmsController{},"get:MobRegister")
}


