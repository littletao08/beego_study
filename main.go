package main

import (
	_"beego_study/routers"
   "github.com/astaxie/beego"
	"beego_study/utils"
)

func main() {

	host := "localhost"
	port := "8080"
	beego.AddFuncMap("zhtime",utils.ZhTime)
	beego.Run(host, port)

//	beego.Run()
//	beego.Error("**************************")

}

