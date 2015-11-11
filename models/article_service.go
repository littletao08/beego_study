package models
import (
	"github.com/astaxie/beego/orm"
	"beego_study/entities"
	"fmt"
)

func init() {
	orm.RegisterModel(new(entities.Article))
}

func Articles(page int) ([]entities.Article, error) {
	var err error
	var articles []entities.Article
	orm := orm.NewOrm()
	_,err = orm.QueryTable("article").All(&articles, "id","user_id", "title", "tag","content","created_at","updated_at")
	return articles, err
}

func LastArticle() (entities.Article, error) {
	var err error
	var article entities.Article
	orm := orm.NewOrm()
	err = orm.QueryTable("article").OrderBy("-id").One(&article, "id","user_id", "title", "tag","content","created_at","updated_at")
	return article, err
}


func Save(article *entities.Article) error{
	var err error
	orm := orm.NewOrm()
	fmt.Printf("****",article)
	_,err=orm.Insert(article)
	return err
}