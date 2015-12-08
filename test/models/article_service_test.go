package models_test
import (
	"testing"
	"time"
	"beego_study/models"
	"beego_study/entities"
	"fmt"
	_"beego_study/initials"
	"github.com/astaxie/beego"
	"beego_study/db"
)



func TestSave(t *testing.T) {
	article := entities.Article{UserId:1, Title:"title", Content:"content", CreatedAt:time.Now()}
	err := models.SaveArticle(&article)
	fmt.Println("err:", err)
}

func TestTransaction(t *testing.T) {
	db := db.NewDB()

	err := db.Begin()
	var sql = "update article set view_count=view_count+1  where user_id = ? and id = ? "
    db.Execute(sql,[]interface{}{1,35120})
	beego.Error("*********************")

	if err == nil {
		db.Commit()
	}else {
		db.Rollback()
	}
}
