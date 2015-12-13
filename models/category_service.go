package models
import (
	"beego_study/entities"
	_"github.com/astaxie/beego/orm"
	"beego_study/constants"
	"beego_study/utils/redis"
	"beego_study/db"
	"bytes"
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


func BatchSaveOrUpdateCategory(db db.DB,categories []entities.Category) {
	sql := bytes.NewBufferString("insert into category(user_id,name,`order`,article_count,created_at) ")
	sql.WriteString("values(?,?,0,1,now()) ")
	sql.WriteString("on duplicate key update article_count =article_count+1,updated_at=now()")
	for _, category := range categories {
		db.Execute(sql.String(), []interface{}{category.UserId, category.Name})
	}

}



