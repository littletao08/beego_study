package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("", &controllers.IndexController{}, "get:Index")

	beego.Router("/users/login", &controllers.UserController{}, "get:Login")
	beego.Router("/users/oauth_login", &controllers.UserController{}, "get:OauthLogin")
	beego.Router("/users/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/users/session", &controllers.UserController{}, "post:Session")
	beego.Router("/users/oauth_session", &controllers.UserController{}, "post:OauthSession")
	beego.Router("/users/register", &controllers.UserController{}, "get:Register")
	beego.Router("/users/oauth_register", &controllers.UserController{}, "get:OauthRegister")
	beego.Router("/users/register_captcha", &controllers.UserController{}, "post:CreateRegisterCaptcha")
	beego.Router("/users/get_test", &controllers.UserController{}, "get:GetTest")
	beego.Router("/users/post_test", &controllers.UserController{}, "post:PostTest")



	beego.Router("/users", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/users/:name", &controllers.UserController{}, "get:UserHome")

	beego.Router("/users/check_user_name", &controllers.UserController{}, "post:CheckUserName")
	beego.Router("/users/check_user_mail", &controllers.UserController{}, "post:CheckUserMail")

	beego.Router("/articles", &controllers.ArticleController{}, "get:Articles")
	beego.Router("/articles", &controllers.ArticleController{}, "post:CreateArticle")
	beego.Router("/articles/new", &controllers.ArticleController{}, "get:New")
	beego.Router("/articles/:id/edit", &controllers.ArticleController{}, "get:EditArticle")
	/*beego.Router("/articles/:id([0-9]+)", &controllers.ArticleController{}, "get:ArticleDetail")*/
	beego.Router("/users/:userId([0-9]+)/articles/:id([0-9]+)", &controllers.ArticleController{}, "get:ArticleDetail")
	beego.Router("/articles", &controllers.ArticleController{}, "put:UpdateArticle")
	beego.Router("/articles/:id([0-9]+)/likes", &controllers.ArticleController{}, "post:Like")

	beego.Router("/categories/:category", &controllers.ArticleController{}, "get:ArticlesGyCategory")
	beego.Router("/users/:userId([0-9]+)/categories/:category", &controllers.ArticleController{}, "get:ArticlesGyUserIdAndCategory")

	beego.Router("/sponsors/new", &controllers.SponsorController{}, "post:New")

	beego.Router("/open_users/:type/auth", &controllers.OpenUserController{}, "get:QqAuth")
	beego.Router("/open_users/:type/token", &controllers.OpenUserController{}, "get:QqToken")

	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = ctx.Input.Query("_method")
		}
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers","Content-Type")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
}


