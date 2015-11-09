package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	_"beego_study/initials"
)
func init() {
	beego.Router("/", &controllers.IndexController{},"get:Index")
	beego.Router("/login", &controllers.UserController{},"get:Login")
	//beego.Router("/catetories", &controllers.CategoryController{},"get:Categories")
}
