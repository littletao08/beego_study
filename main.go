package main

import (
	_"beego_study/routers"
	"github.com/astaxie/beego"
"beego_study/utils"
)

func main() {

	/*host := "192.168.10.43"
	port := "8080"
	beego.Run(host, port)*/

	beego.AddFuncMap("zhtime", utils.ZhTime)
	beego.Run()

}

