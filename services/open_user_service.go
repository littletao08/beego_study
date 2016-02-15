package services

import (
	"github.com/astaxie/beego/orm"
	"beego_study/entities"
	"bytes"
	"beego_study/db"
)


func OpenUser(openId string, openType int) (*entities.OpenUser, error) {
	var err error
	var openUser entities.OpenUser

	orm := orm.NewOrm()
	err = orm.QueryTable("open_user").Filter("open_id", openId).Filter("type", openType).One(&openUser)

	return &openUser, err
}


func SaveOrUpdateOpenUser(openUser *entities.OpenUser) (error) {
	sql := bytes.NewBufferString("insert ignore into open_user(open_id,user_id,type,nick,head,sex,age,province,city,created_at) ")
	sql.WriteString("values(?,?,?,?,?,?,?,?,?,now()) ")
	sql.WriteString("on duplicate key update nick =?,head=?,sex=?,province=?,city=?")
	db := db.NewDB()
	params:=[]interface{}{openUser.OpenId, openUser.UserId, openUser.Type, openUser.Nick}
	params=append(params,openUser.Head,openUser.Sex, openUser.Age, openUser.Province, openUser.City)
	params=append(params,openUser.Nick,openUser.Head,openUser.Sex,openUser.Province,openUser.City)
	_ , err := db.Execute(sql.String(), params)
	return err
}

func BindUserIdToOpenUser(openId string, openType int, userId int64) (int64, error) {
	sql := "update open_user set user_id = ? where open_id = ? and type = ? "
	db := db.NewDB()
	result, err := db.Execute(sql, []interface{}{userId, openId, openType})
	return result, err
}


