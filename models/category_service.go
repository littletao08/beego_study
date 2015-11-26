package models
import (
	"beego_study/entities"
	"github.com/astaxie/beego/orm"
	"beego_study/constants"
	"beego_study/utils/redis"
)

func Categories() ([]entities.Category, error) {
	var err error
	var categories []entities.Category
	var categoriesKey = constants.CATEGORY_KEY
	err = redis_util.Get(categoriesKey, &categories)
	if err == nil {
		return categories, nil;
	}
	orm := orm.NewOrm()
	_, err = orm.QueryTable("category").OrderBy("order").All(&categories, "id", "name", "order", "created_at", "updated_at")
	return categories, err
}