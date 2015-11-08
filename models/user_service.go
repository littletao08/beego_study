package models
import (
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"beego_study/constants"
	"beego_study/utils"
	"fmt"
)

func init() {
	orm.RegisterModel(new(entities.User))
}

func User(id int) (entities.User, error) {
	var err error
	var user entities.User
	var userKey = constants.USER_KEY + string(id);
	err = utils.Get(userKey, &user)
	fmt.Println("************err:", err, "user:", user)
	if err == nil {
		return user, nil;
	}
	orm := orm.NewOrm()
	err = orm.QueryTable("user").Filter("id", id).One(&user, "id", "nick", "age", "cell","mail","sex", "CreatedAt","UpdatedAt")
	if err == nil {
		utils.Set(userKey, user, 1000)
	}
	return user, err
}