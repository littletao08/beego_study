package models

import (
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"beego_study/exception"
	"strconv"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"time"
)

func User(id int64) (entities.User, error) {
	var err error
	var user entities.User
	/*var userKey = constants.USER_KEY + string(id);
	err = redis_util.Get(userKey, &user)
	if err == nil {
		return user, nil;
	}*/
	orm := orm.NewOrm()
	err = orm.QueryTable("user").Filter("id", id).One(&user, "id","name", "nick", "age", "cell", "mail", "sex", "CreatedAt", "UpdatedAt")
	/*if err == nil {
		redis_util.Set(userKey, user, 1000)
	}*/
	return user, err
}

func SaveUser(user *entities.User) error {
	orm := orm.NewOrm()

	err := CheckNewUser(user)

	if nil != err {
		return err
	}

	user.CreatedAt=time.Now()

	id, err := orm.Insert(user)
	if nil == err {
		return err
	}

	user.Id = id
	return nil
}
func NewMobUser(user *entities.User)error{
	orm := orm.NewOrm()

	err := CheckUserMobile(user.Cell)

	if nil != err {
		return err
	}
	user.CreatedAt=time.Now()

	id, err := orm.Insert(user)
	if nil == err {
		return err
	}
	user.Id = id
	return nil
}

func FundUser(name string, password string) (entities.User, error) {
	var err error
	var user entities.User
	beego.Error("name:", name, "password:", password)

	orm := orm.NewOrm()
	querySetter := orm.QueryTable("user").Filter("password", password);

	valid := validation.Validation{}

	if valid.Email(name, "email").Ok {
		querySetter = querySetter.Filter("mail", name)
	}else {
		querySetter = querySetter.Filter("name", name)
	}

	err = querySetter.One(&user, "id", "name", "nick", "password", "age", "cell", "mail", "sex", "CreatedAt", "UpdatedAt")
	return user, err
}

func CheckUserName(name string) error {

	minNameLength := ParameterIntValue("user_name_min_length")
	maxNameLength := ParameterIntValue("user_name_max_length")

	beego.Error("minNameLength:", minNameLength, "maxNameLength:", maxNameLength)
	nameLength := len(name)
	if nameLength < minNameLength || nameLength > maxNameLength {
		return errors.New("用户名长度只能在" + strconv.Itoa(minNameLength) + "-" + strconv.Itoa(maxNameLength) + "字符之间")
	}

	orm := orm.NewOrm()
	count, err := orm.QueryTable("user").Filter("name", name).Count()

	if nil != err || count > 0 {
		return exception.USER_NAME_EXISTENT
	}
	return nil

}

func CheckUserMail(mail string) error {

	valid := validation.Validation{}

	if !valid.Email(mail, "email").Ok {
		return errors.New("邮箱格式错误")
	}

	orm := orm.NewOrm()

	count, err := orm.QueryTable("user").Filter("mail", mail).Count()

	if nil != err || count > 0 {
		return exception.USER_MAIL_EXISTENT
	}

	return nil
}

func CheckUserMobile(mobile string) error {

	valid := validation.Validation{}

	if !valid.Mobile(mobile, "mobile").Ok {
		return errors.New("手机号格式不正确")
	}

	orm := orm.NewOrm()

	count, err := orm.QueryTable("user").Filter("cell", mobile).Count()

	if nil != err || count > 0 {
		return exception.USER_MOBILE_EXISTENT
	}

	return nil
}

func CheckNewUser(user *entities.User) error {

	name := user.Name
	mail := user.Mail
	pass := user.Password

	minPassLength := ParameterIntValue("user_pass_min_length")
	maxPassLength := ParameterIntValue("user_pass_max_length")
	passLength := len(pass)

	if passLength < minPassLength || passLength > maxPassLength {
		return errors.New("密码长度只能在" + strconv.Itoa(minPassLength) + "-" + strconv.Itoa(maxPassLength) + "字符之间")
	}

	var err = CheckUserName(name)
	if nil != err {
		return err
	}

	err = CheckUserMail(mail)
	if nil != err {
		return err
	}

	return nil
}