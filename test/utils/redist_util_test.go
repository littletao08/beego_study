package utils_test
import (
	"beego_study/utils"
	"fmt"
	"github.com/astaxie/beego"
	_ "beego_study/routers"
	"testing"
)

func init() {
	beego.AppConfig.DefaultString("cache", "redis")
}

func TestSet(t *testing.T) {
	err := utils.Set("aa", "aa", 10);
	fmt.Print("val:", err)
}

func TestGet(t *testing.T) {
	var aa string
	utils.Get("aa",aa);
	fmt.Print("val:", aa)
}

