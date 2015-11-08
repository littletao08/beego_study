package routers

import (
	"github.com/astaxie/beego"
	"beego_study/controllers"
	_"beego_study/initials"
)
func init() {
	beego.Router("/", &controllers.IndexController{},"get:Index")
	//beego.Router("/users/:id", &controllers.UserController{},"get:Users")
	//beego.Router("/catetories", &controllers.CategoryController{},"get:Categories")
}



