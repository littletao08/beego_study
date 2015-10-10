package routers


import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
)
func init() {
	beego.Router("/", &controllers.IndexController{},"get:Index")
	beego.Router("/users/:id", &controllers.UserController{},"get:Users")
}



