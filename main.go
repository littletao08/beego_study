package main

import (
	"github.com/astaxie/beego"
	_"beego_study/routers"

)


func main() {
	host := "192.168.10.147"
	port := "8080"
	beego.Run(host, port)
}

