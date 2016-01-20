package utils_test

import (
	"testing"
	"beego_study/entities"
	"beego_study/utils"
	"fmt"
	_"beego_study/initials"
	"github.com/astaxie/beego/validation"
)

func TestExtractFieldValues(t *testing.T) {

	var users []entities.User

	users = append(users, entities.User{Id:1, Name:"第1个"})
	users = append(users, entities.User{Id:2, Name:"第2个"})
	users = append(users, entities.User{Id:3, Name:"第3个"})
	users = append(users, entities.User{Id:4, Name:"第4个"})
	users = append(users, entities.User{Id:5, Name:"第5个"})
	users = append(users, entities.User{Id:6, Name:"第6个"})

	ids, err := utils.ExtractFieldValues(users, "Id")
	if nil != err {
		fmt.Println("err:", err)
	}
	for _, id := range ids {
		fmt.Println("id:", id)
	}

	str := utils.SliceToString(ids, ",")

	fmt.Println("str:", str)

}

func TestEmailValidate(t *testing.T) {
	mail := "406504302qq.com"
	valid := validation.Validation{}

	if !valid.Email(mail, "email").Ok {
		fmt.Println("mail 格式错误")
	}
}

