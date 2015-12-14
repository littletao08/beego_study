package models_test
import (
	"fmt"
	"beego_study/models"
	"testing"
	_ "github.com/astaxie/beego"
	_"beego_study/test/initials"
)

func TestGetUsers(t *testing.T) {
	fmt.Println("*************************")
	var v, _ = models.ParameterValue("x-bmob-application-id")
	fmt.Println("v:", v)
}
