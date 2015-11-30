package utils_test
import (
	"github.com/astaxie/beego"
	_ "beego_study/routers"
	"beego_study/initials"
)

func init() {
	beego.AppConfig.DefaultString("cache", "redis")
	initials.InitRedis()
}

