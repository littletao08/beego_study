package models
import (
	"github.com/astaxie/beego/orm"
	"beego_study/entities"
	"beego_study/db"
)

func Articles(page int) ([]entities.Article, error) {
	var err error
	var articles []entities.Article
	db := db.NewDB()
	_, err = db.QueryTable("article").All(&articles, "id", "user_id", "title", "tag", "content", "created_at", "updated_at")
	return articles, err
}

func AllArticles(userId int64, pagination *db.Pagination) {
	db := db.NewDB()
	var articles []entities.Article
	db.From("article").OrderBy("created_at desc").FillPagination(&articles, pagination)
}

func LastArticle() (entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").OrderBy("-id").One(&article, "id", "user_id", "title", "tag", "content", "created_at", "updated_at")
	return article, err
}

func ArticleByIdAndUserId(userId int64,articleId int64) (*entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").Filter("user_id",userId).Filter("id",articleId).One(&article, "id", "user_id", "title", "tag", "content", "created_at", "updated_at")
	return &article, err
}

func ArticleById(articleId int64) (*entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").Filter("id",articleId).One(&article, "id", "user_id", "title", "tag", "content", "created_at", "updated_at")
	return &article, err
}

func Save(article *entities.Article) error {
	var err error
	orm := orm.NewOrm()
	_, err = orm.Insert(article)

	return err
}


