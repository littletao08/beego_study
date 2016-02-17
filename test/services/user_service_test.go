package services

import (
	_"beego_study/test/initials"
	"testing"
	"beego_study/entities"
	"beego_study/services"
	"fmt"
)

func TestSaveUser(t *testing.T){

	user := new(entities.User)
	user.Name="name"
	user.Mail="mail@qq.com"
	user.Password="password"

	err := services.SaveUser(user)

	fmt.Println(err)
}
