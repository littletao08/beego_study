package models_test
import (
	"testing"
	"time"
	"beego_study/entities"
	"fmt"
	_"beego_study/test/initials"
	_"beego_study/initials"
	"beego_study/db"
	"beego_study/models"
	"html"
	"strings"
	"github.com/goquery"
	"github.com/astaxie/beego"
	"github.com/gogather/com"
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

func TestGet(t *testing.T) {
	article, _ := models.ArticleById(35120)
	var content = html.UnescapeString(article.Content)
	if len(content) == 0 {
		return
	}
	reader := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(reader)
	text := doc.Text()
	text = strings.Replace(text," ","",-1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)
	subText := com.SubString(text, 0, 160)
	beego.Error(subText)
}
