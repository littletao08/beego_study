package services

import (
	_ "beego_study/test/initials"
	"testing"
	"beego_study/services"
	"beego_study/entities"
	"fmt"
)

func TestOpenUser(t *testing.T)  {
	openuser , err := services.OpenUser("1",entities.OPEN_USER_TYPE_QQ)

	fmt.Println("openuser",openuser,"err",err)
}

func TestSaveOrUpdateOpenUser(t *testing.T) {
	openUser := new(entities.OpenUser)
	openUser.Age=26
	openUser.City="beijing"
	openUser.Head="www.baidu.com/header.png"
	openUser.Nick="nick"
	openUser.OpenId="openid"
	openUser.Province="province"
	openUser.Sex=2
	openUser.Type=entities.OPEN_USER_TYPE_QQ
	services.SaveOrUpdateOpenUser(openUser)
}


