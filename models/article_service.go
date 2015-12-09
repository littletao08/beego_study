package models
import (
	"github.com/astaxie/beego/orm"
	"beego_study/entities"
	"beego_study/db"
	"time"
	"errors"
	"beego_study/exception"
	"beego_study/utils"
)

func Articles(page int) ([]entities.Article, error) {
	var err error
	var articles []entities.Article
	db := db.NewDB()
	_, err = db.QueryTable("article").All(&articles, "id", "user_id", "title", "tag", "content", "view_count", "like_count", "comment_count", "created_at", "updated_at")

	return articles, err
}

func AllArticles(userId int64, pagination *db.Pagination) {
	db := db.NewDB()
	var articles []entities.Article
	db.From("article").OrderBy("created_at desc").FillPagination(&articles, pagination)

	//设置当前用户点赞标记
	if len(pagination.Data) < 1 || userId <= 0 {
		return
	}

	ids, err := utils.ExtractFieldValues(pagination.Data, "Id")
	if nil != err {
		return
	}
	signs := ArticleLikeSigns(userId, ids)
	signMap := make(map[int64]bool)
	for _, v := range signs {
		signMap[v] = true
	}
	var newArticles []entities.Article
	for _, v := range articles {
		v.HasLike = signMap[v.Id]
		newArticles = append(newArticles, v)
	}

	pagination.SetData(newArticles)
}

func LastArticle() (entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").OrderBy("-id").One(&article, "id", "user_id", "title", "tag", "content", "view_count", "like_count", "comment_count", "created_at", "updated_at")
	return article, err
}

func ArticleByIdAndUserId(userId int64, articleId int64) (*entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").Filter("user_id", userId).Filter("id", articleId).One(&article, "id", "user_id", "title", "tag", "content", "view_count", "like_count", "comment_count", "created_at", "updated_at")
	return &article, err
}

func ArticleById(articleId int64) (*entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").Filter("id", articleId).One(&article, "id", "user_id", "title", "tag", "content", "view_count", "like_count", "comment_count", "created_at", "updated_at")
	return &article, err
}

func SaveArticle(article *entities.Article) error {
	var err error
	orm := orm.NewOrm()
	_, err = orm.Insert(article)

	return err
}


func IncrViewCount(articleId int64, userId int64, ip string) error {

	db := db.NewDB()
	err := db.Begin()

	hasViewed, err := HasViewArticle(articleId, userId, ip, db)
	if nil != err || hasViewed {
		db.Rollback()
		return err
	}
	article, err := ArticleById(articleId)

	articleOwnerId := article.UserId
	if err != nil {
		db.Rollback()
		return err
	}

	if articleOwnerId <= 0 {
		db.Rollback()
		return errors.New(exception.NOT_EXIST_ARTICLE_ERROR.Message())
	}

	var sql = "update article set view_count=view_count+1  where user_id = ? and id = ? "

	_, err = db.Execute(sql, []interface{}{articleOwnerId, articleId})
	if nil == err {
		var articleView = entities.ArticleView{UserId:userId, ArticleId:articleId, Ip:ip, CreatedAt:time.Now()}
		err = SaveArticleView(articleView, db)
	}

	if nil == err {
		err = db.Commit()
	}else {
		err = db.Rollback()
	}

	return err
}


func IncrLikeCount(articleId int64, userId int64) (int, error) {
	db := db.NewDB()
	err := db.Begin()

	hasLiked, err := HasLikeArticle(articleId, userId, db)
	if nil != err {
		db.Rollback()
		return 0, err
	}

	article, err := ArticleById(articleId)

	articleOwnerId := article.UserId
	if err != nil {
		db.Rollback()
		return 0, err
	}

	if articleOwnerId <= 0 {
		db.Rollback()
		return 0, errors.New(exception.NOT_EXIST_ARTICLE_ERROR.Message())
	}

	var sql = "update article set like_count=like_count+? where user_id = ? and id = ? "
	var incrCount, valid int = 1, 1
	if hasLiked {
		incrCount = -1
		valid = 0
	}
	_, err = db.Execute(sql, []interface{}{incrCount, articleOwnerId, articleId})
	if nil == err {
		var articleLike = entities.ArticleLike{UserId:userId, ArticleId:articleId, Valid:valid, CreatedAt:time.Now()}
		err = SaveOrUpdate(articleLike, db)
	}

	if nil == err {
		err = db.Commit()
	}else {
		incrCount = 0
		err = db.Rollback()
	}

	return incrCount, err
}


