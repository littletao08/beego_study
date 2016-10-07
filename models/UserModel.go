package models

import "beego_study/entities"

type UserModel struct {
	entities.User
}

func (m UserModel) AsJSON() (map[string]interface{}){
	var r = Slice(m, "Name","Age");
	return r
}
