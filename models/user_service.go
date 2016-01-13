package models
import (
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"beego_study/constants"
	"beego_study/utils/redis"
	"time"
)

func User(id int) (entities.User, error) {
	var err error
	var user entities.User
	var userKey = constants.USER_KEY + string(id);
	err = redis_util.Get(userKey, &user)
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

func CreateNewUser(name string, password string) (entities.User,error) {
	var user entities.User
	user.Name = name
	user.Password = password
	orm := orm.NewOrm()
	var err error
	_, err = orm.Insert(&user)
	return user, err;
}

func CreateQCUser(name string, qcOpenId string,sex string) (*entities.User,error) {
	var user entities.User
	user.Name = name
	user.Nick = name
	user.QcOpenId = qcOpenId
	user.Sex = 1
	user.CreatedAt=time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt=time.Now().Format("2006-01-02 15:04:05")

	orm := orm.NewOrm()
	var err error
	_, err = orm.Insert(&user)
	return &user, err;
}


func SetUserPassword(user *entities.User) error {
	orm :=orm.NewOrm()
	_,err := orm.Update(user);
	return err
}


func FundQCUser(qcOpenId string)(*entities.User,error){
	var err error
	var user entities.User
	orm :=orm.NewOrm()
	err =orm.QueryTable("user").Filter("qc_open_id",qcOpenId).One(&user,"id", "name", "nick", "password", "age", "cell", "mail", "sex", "QcOpenId","CreatedAt", "UpdatedAt")

	return &user, err
}