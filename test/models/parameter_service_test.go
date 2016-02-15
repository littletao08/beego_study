package services
import (
	"fmt"
	"beego_study/services"
	"testing"
	_ "github.com/astaxie/beego"
	_"beego_study/test/initials"
	"github.com/astaxie/beego"
	"beego_study/entities"
)

func TestGetUsers(t *testing.T) {
	fmt.Println("*************************")
	var v, _ = services.ParameterValue("x-bmob-application-id")
	fmt.Println("v:", v)
}



func TestToLower(t *testing.T){
	artile := entities.Article{}
	prt := &artile
	prt.SetTags("Artile")
	beego.Error(artile.Tags)
}