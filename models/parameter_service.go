package models
import (
	"beego_study/constants"
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"time"
	"beego_study/utils/redis"
)

func ParameterValue(key string) (interface{}, error) {

	var v interface{}
	var err error
	var parameter entities.Parameter
	var parametersKey = constants.PARAMETERS_KEY
	err = redis_util.Hget(parametersKey, key, &parameter)
	if nil == err {
		return parameter.Value, err
	}

	db := orm.NewOrm()
	err = db.QueryTable("parameter").Filter("key", key).One(&parameter)
	if nil == err {
		redis_util.Hset(parametersKey, key, parameter, int64(time.Hour * 24 / time.Second))
		v = parameter.Value
	}
	return v, err
}
