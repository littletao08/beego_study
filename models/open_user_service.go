package models

import (
	"github.com/astaxie/beego/orm"
	"beego_study/entities"
	"bytes"
	"beego_study/db"
)


func Oauth(openId string, openType int) (entities.OpenUser, error) {
	var err error
	var oauth entities.OpenUser

	orm := orm.NewOrm()
	err = orm.QueryTable("oauth").Filter("open_id", openId).Filter("type", openType).One(&oauth, "open_id", "user_id", "nick", "sex", "age")

	return oauth, err
}

func SaveOauth(oauth entities.OpenUser) (int, error) {
	sql := bytes.NewBufferString("insert ignore into oauth(open_id,user_id,type,nick,sex,age,province,city,created_at) ")
	sql.WriteString("values(?,?,?,?,?,?,?,now()) ")
	db := db.NewDB()
	result, err := db.Execute(sql.String(), []interface{}{oauth.OpenId, oauth.UserId, oauth.Type, oauth.Nick, oauth.Sex, oauth.Age, oauth.Province, oauth.City})
	return result, err
}

func BindUserIdToOpenUser(openId string, openType int, userId int64) (int, error) {
	sql := "update oauth set user_id = ? where open_id = ? and type = ? "
	db := db.NewDB()
	result, err := db.Execute(sql, []interface{}{userId, openId, openType})
	return result, err
}


