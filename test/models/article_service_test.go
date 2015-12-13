package models_test
import (
	"testing"
	"time"
	"beego_study/entities"
	"fmt"
	_"beego_study/initials"
	"beego_study/db"
	"beego_study/models"
)



func TestSave(t *testing.T) {
	article := entities.Article{UserId:1, Title:"title", Tags:"go,reis", Categories:"go3,go4", Content:"content", CreatedAt:time.Now()}
	models.SaveArticle(&article)
	db := db.NewDB()
	db.Insert(&article)

	fmt.Println("&article:", &article)
}

func TestTransaction(t *testing.T) {
	db := db.NewDB()

	err := db.Begin()
	var sql = "update article set view_count=view_count+1  where user_id = ? and id = ? "
	db.Execute(sql, []interface{}{1, 35120})

	if err == nil {
		db.Commit()
	}else {
		db.Rollback()
	}
}

func TestRel(t *testing.T) {
	db := db.NewDB()
	var articles []entities.Article
	db.QueryTable("article").RelatedSel().Filter("user_id", 1).All(&articles)
}
