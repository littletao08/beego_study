package services

import (
	"beego_study/entities"
	"beego_study/db"
	"time"
	"errors"
	"beego_study/exception"
	"beego_study/utils"
	"bytes"
	"strings"
	"github.com/astaxie/beego"
)


func AllArticles(userId int64, pagination *db.Pagination) {
	db := db.NewDB()
	var articles []entities.Article
	query := db.From("article")

	if userId > 0 {
		query.Where("user_id", userId)
	}

	query.OrderBy("created_at desc").FillPagination(&articles, pagination)

}

func SetLikeSign(pagination *db.Pagination, userId int64) {
	//设置当前用户点赞标记
	if len(pagination.Data) < 1 || userId <= 0 {
		return
	}

	ids, err := utils.ExtractFieldValues(pagination.Data, "Id")
	if nil != err {
		return
	}

	signs := ArticleLikeLogs(userId, ids)
	signMap := make(map[int64]bool)
	for _, v := range signs {
		signMap[v] = true
	}

	var articles []entities.Article
	for _, v := range pagination.Data {
		article := v.(entities.Article)
		article.HasLike = signMap[article.Id]
		articles = append(articles,article)
	}
	pagination.SetData(articles)
}

func ArticlesGyCategory(userId int64, category string, pagination *db.Pagination,isFilterUserId bool) {
	db := db.NewDB()
	var articles []entities.Article
	query := db.From("article")
	if userId > 0 && isFilterUserId {
		query.Where("user_id", userId)
	}

	if len(category) > 0 {
		category = strings.ToLower(category)
		beego.Error("category", category)
		query.Like("categories", "%" + category + "%")
	}

	query.OrderBy("created_at desc").FillPagination(&articles, pagination)

	//设置当前用户点赞标记
	if len(pagination.Data) < 1 || userId <= 0 {
		return
	}

	ids, err := utils.ExtractFieldValues(pagination.Data, "Id")
	if nil != err {
		return
	}
	if userId > 0 {
		signs := ArticleLikeLogs(userId, ids)
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
}

func LastArticle() (entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").OrderBy("-id").One(&article)
	return article, err
}

func ArticleByIdAndUserId(articleId int64, userId int64) (*entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").Filter("user_id", userId).Filter("id", articleId).One(&article)
	return &article, err
}

func ArticleById(articleId int64) (*entities.Article, error) {
	var err error
	var article entities.Article
	db := db.NewDB()
	err = db.QueryTable("article").Filter("id", articleId).One(&article)

	return &article, err
}

func SaveArticle(article *entities.Article) error {

	var err error
	db := db.NewDB()
	db.Begin()

	bBuffer := bytes.NewBufferString("insert into article (user_id,title, tags,categories, content, created_at) ")
	bBuffer.WriteString("values(?,?,?,?,?,now())")

	_, err = db.Raw(bBuffer.String(), []interface{}{article.UserId, article.Title, article.Tags, article.Categories, article.Content}).Exec()

	if nil == err {
		var categories []entities.Category
		if len(article.Categories) > 0 {
			categoryNames := strings.Split(article.Categories, ",")
			categories = entities.NewCategories(article.UserId, categoryNames)
		}
		BatchSaveOrUpdateCategory(db, categories)
	}

	if nil == err {
		db.Commit()
	}else {
		db.Rollback()
	}

	return err
}

func UpdateArticle(article *entities.Article) error {

	var err error
	db := db.NewDB()
	db.Begin()

	sql := "update article set title = ? ,tags=?,categories=?, content=?, updated_at=now() where user_id = ? and  id = ? "

	_, err = db.Raw(sql, []interface{}{article.Title, article.Tags, article.Categories, article.Content, article.UserId, article.Id}).Exec()

	if nil == err {
		var categories []entities.Category
		if len(article.Categories) > 0 {
			categoryNames := strings.Split(article.Categories, ",")
			categories = entities.NewCategories(article.UserId, categoryNames)
		}
		BatchSaveOrUpdateCategory(db, categories)
	}

	if nil == err {
		db.Commit()
	}else {
		db.Rollback()
	}

	return err
}

func IncrViewCount(articleId int64, userId int64, ip string) (bool, error) {

	db := db.NewDB()
	err := db.Begin()

	hasViewed, err := HasViewArticle(articleId, userId, ip, db)
	if nil != err || hasViewed {
		db.Rollback()
		return false, err
	}
	article, err := ArticleById(articleId)

	articleOwnerId := article.UserId
	if err != nil {
		db.Rollback()
		return false, err
	}

	if articleOwnerId <= 0 {
		db.Rollback()
		return false, errors.New(exception.NOT_EXIST_ARTICLE_ERROR.Error())
	}

	var sql = "update article set view_count=view_count+1  where user_id = ? and id = ? "

	_, err = db.Execute(sql, []interface{}{articleOwnerId, articleId})
	if nil == err {
		var articleView = entities.ArticleView{UserId:userId, ArticleId:articleId, Ip:ip, CreatedAt:time.Now()}
		err = SaveArticleView(articleView, db)
	}

	if nil == err {
		sql = "update user set view_count=view_count+1  where id = ? "
		_, err = db.Execute(sql, []interface{}{articleOwnerId})
	}

	if nil == err {
		err = db.Commit()
	}else {
		err = db.Rollback()
	}

	return nil == err, err
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
		return 0, errors.New(exception.NOT_EXIST_ARTICLE_ERROR.Error())
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
		sql = "update user set like_count=like_count+1  where id = ? "
		_, err = db.Execute(sql, []interface{}{articleOwnerId})
	}

	if nil == err {
		err = db.Commit()
	}else {
		incrCount = 0
		err = db.Rollback()
	}

	return incrCount, err
}

func ValidArticle(a entities.Article) (error, bool) {

	var title = a.Title
	var category = a.Categories;
	var tag = a.Tags;
	var content = a.Content

	var titleMaxLen = ParameterIntValue("article_title_max_length")
	var categoryMaxLen = ParameterIntValue("article_categories_max_length")
	var tagMaxLen = ParameterIntValue("article_tags_max_length")
	var contentMaxLen = ParameterIntValue("article_content_max_length")

	if len(title) > titleMaxLen {
		return exception.ARTICLE_TITLE_LEN_OVERFLOW, false
	}

	if len(category) > categoryMaxLen {
		return exception.ARTICLE_CATEGORY_LEN_OVERFLOW, false
	}

	if len(tag) > tagMaxLen {
		return exception.ARTICLE_TAG_LEN_OVERFLOW, false
	}

	if len(content) > contentMaxLen {
		return exception.ARTICLE_CONTENT_LEN_OVERFLOW, false
	}

	return nil, true

}

func TopLikeArticles() []entities.Article  {

    var articles []entities.Article
	sql := "select id,user_id,title from article where like_count >0 and  created_at > date_sub(now(),interval ? DAY) order by like_count desc limit 10"

	dayRange := ParameterIntValue("like_article_day_range")

	db := db.NewDB()
	db.Raw(sql,[]interface{}{dayRange}).QueryRows(&articles)
	return articles ;
}



