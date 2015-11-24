package test_test
import (
"fmt"
"beego_study/models"
"testing"
"github.com/astaxie/beego"
"beego_study/initials"
)
func init()  {
		beego.AppConfig.DefaultString("cache", "redis")
		initials.InitRedis()

}
func TestGetUsers(t *testing.T) {

	fmt.Println("*************************")
	var v,_  = models.ParameterValue("x-bmob-application-id")
	fmt.Println("v:",v)
}
