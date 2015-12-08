package models
import (
	"beego_study/entities"
	"beego_study/db"
	"bytes"
)

func SaveOrUpdate(articleLike entities.ArticleLike, db db.DB) error {
	sql := bytes.NewBufferString("insert into article_like (user_id , article_id,valid,created_at) ")
	sql.WriteString("values(?,?,?,now()) ")
	sql.WriteString("on duplicate key update valid =?,updated_at=now() ")
	_, err := db.Execute(sql.String(), []interface{}{articleLike.UserId, articleLike.ArticleId, articleLike.Valid, articleLike.Valid})
	return err

}


func HasLikeArticle(articleId int64, userId int64, db db.DB) (bool, error) {
	count, err := db.QueryTable("article_like").Filter("user_id", userId).Filter("article_id", articleId).Filter("valid", 1).Count()

	return count > 0, err
}

