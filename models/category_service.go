package models
import (
	"beego_study/entities"
	_"github.com/astaxie/beego/orm"
	"beego_study/constants"
	"beego_study/utils/redis"
	"beego_study/db"
)

func Categories() ([]entities.Category, error) {
	var err error
	var categories []entities.Category
	var categoriesKey = constants.CATEGORY_KEY
	err = redis_util.Get(categoriesKey, &categories)
	if err == nil {
		return categories, nil;
	}
	db := db.NewDB()
	_, err = db.QueryTable("category").OrderBy("order").All(&categories, "id", "name", "order", "created_at", "updated_at")
	return categories, err
}