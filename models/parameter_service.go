package models
import (
	"beego_study/constants"
	"beego_study/utils"
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
)
func init() {
	orm.RegisterModel(new(entities.Parameter))
}

func ParameterValue(key string) (interface{}, error) {

	var v interface{}
	var err error
	var parameterMap = make(map[string]entities.Parameter)
	var parametersKey = constants.PARAMETERS_KEY
	err = utils.Get(parametersKey, &parameterMap)
	if nil == err {
		var parameter = parameterMap[key]
		return parameter.Value, err
	}

	var parameter entities.Parameter
	orm := orm.NewOrm()
	err = orm.QueryTable("parameter").Filter("key", key).One(&parameter)
	if nil == err {
		parameterMap[key] = parameter
		v = parameter.Value
	}
	return v, err
}
