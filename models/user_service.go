package models
import (
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"beego_study/constants"
	"fmt"
	"beego_study/utils/redis"
)

func User(id int) (entities.User, error) {
	var err error
	var user entities.User
	var userKey = constants.USER_KEY + string(id);
	err = redis_util.Get(userKey, &user)
	fmt.Println("************err:", err, "user:", user)
	if err == nil {
		return user, nil;
	}
	orm := orm.NewOrm()
	err = orm.QueryTable("user").Filter("id", id).One(&user, "id", "nick", "age", "cell", "mail", "sex", "CreatedAt", "UpdatedAt")
	if err == nil {
		redis_util.Set(userKey, user, 1000)
	}
	return user, err
}

func FundUser(name string, password string) (entities.User, error) {
	var err error
	var user entities.User
	orm := orm.NewOrm()
	err = orm.QueryTable("user").Filter("name", name).Filter("password", password).One(&user, "id", "name", "nick", "password", "age", "cell", "mail", "sex", "CreatedAt", "UpdatedAt")

	return user, err
}