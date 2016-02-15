package services
import (
	"beego_study/entities"
	"beego_study/db"
)

func SaveArticleView(articleView entities.ArticleView,db db.DB ) error {
	_, err := db.Insert(&articleView)
	return err

}

func HasViewArticle(articleId int64 ,userId int64, ip string,db db.DB) (bool, error) {

	count,err := db.QueryTable("article_view").Filter("article_id",articleId).Filter("ip", ip).Count()

	return count>0,err
}
