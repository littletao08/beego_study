package main

import (
	"github.com/astaxie/beego"
	_"beego_study/routers"
	_"beego_study/initials"
)


func main() {
	host := "localhost"
	port := "8080"
	beego.Run(host, port)
}

