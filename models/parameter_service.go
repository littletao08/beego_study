package models
import (
	"beego_study/constants"
	"beego_study/utils"
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"time"
)
func init() {
	orm.RegisterModel(new(entities.Parameter))
}

func ParameterValue(key string) (interface{}, error) {

	var v interface{}
	var err error
	var parameter entities.Parameter
	var parametersKey = constants.PARAMETERS_KEY
	err = utils.Hget(parametersKey,key, &parameter)
	if nil == err {
		return parameter.Value, err
	}

	orm := orm.NewOrm()
	err = orm.QueryTable("parameter").Filter("key", key).One(&parameter)
	if nil == err {
		utils.HSet(parametersKey,key,parameter,int64(time.Hour*24/time.Second))
		v = parameter.Value
	}
	return v, err
}
