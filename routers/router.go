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


