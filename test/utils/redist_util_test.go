package utils_test
import (
	"beego_study/utils"
	"fmt"
	"github.com/astaxie/beego"
	_ "beego_study/routers"
	"testing"
	"beego_study/initials"
)

func init() {
	beego.AppConfig.DefaultString("cache", "redis")
	initials.InitRedis()
}

