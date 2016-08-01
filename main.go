package main

import (
	_"beego_study/initials"
	_"beego_study/routers"
	"github.com/astaxie/beego"
	"beego_study/utils"
)

func init() {
	beego.BeeLogger.Async()
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogger("file", `{"filename":"logs/beego.log"}`)
}
func main() {

	/*host := "192.168.10.43"
	port := "8080"R
	beego.Run(host, port)*/


	beego.AddFuncMap("zhtime", utils.ZhTime)
	beego.Run()

}
