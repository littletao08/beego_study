package utils_test

import (
	"testing"
	"beego_study/entities"
	"beego_study/models"
	"fmt"
)

func Test(t *testing.T) {
	var user = entities.User{};
	user.Name="张利达"
	user.Age=27
	user.Sex=1
	r := models.Slice(user,"Name","Age")

	fmt.Println("-------",r)

	var userModel = models.UserModel{}
	userModel.Name="张利达"
	userModel.Age=27
	userModel.Sex=1
	var i interface{} = userModel
	_,ok :=i.(models.IModel)
	fmt.Println("--------------",ok)

	fmt.Println("-------",userModel.AsJSON())

}
