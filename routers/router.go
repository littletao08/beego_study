package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	_"beego_study/initials"
	"github.com/astaxie/beego/context"
	"beego_study/entities"
)
func init() {
	beego.Router("", &controllers.IndexController{},"get:Index")

	beego.Router("/login", &controllers.UserController{},"get:Login")

	beego.Router("/users/session", &controllers.UserController{},"post:Session")
	beego.Router("/articles", &controllers.ArticleController{},"post:CreateArticle")
	beego.Router("/articles/new", &controllers.ArticleController{},"get:New")
	beego.Router("/articles/:id([0-9]", &controllers.ArticleController{},"get:ArticleDetail")

	beego.Router("/sponsors/new", &controllers.SponsorController{},"post:New")
	beego.Router("/article/praise/:id([0-9]",&controllers.ArticleController{},"get:ArticlePraise")


	beego.InsertFilter("/*",beego.BeforeRouter,func(ctx *context.Context) {
		_, ok := ctx.Input.Session("user").(entities.User)
		if !ok {
			ctx.Redirect(302, "/login")
		}
	})
}


