package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beego_study/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_study/controllers:MainController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_study/controllers:MainController"] = append(beego.GlobalControllerRouter["beego_study/controllers:MainController"],
		beego.ControllerComments{
			"Users",
			`/users`,
			[]string{"get"},
			nil})

}
