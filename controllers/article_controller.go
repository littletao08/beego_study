package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"time"
	"github.com/astaxie/beego"
	"beego_study/exception"
	"beego_study/db"
)

type ArticleController struct {
	BaseController
}
func (c *ArticleController) Articles() {
	pagination := c.NewPagination()
	userId := c.UserId()
	models.AllArticles(userId, pagination)
	c.Data["pagination"] = pagination
	c.TplName = "index.html"

}

func (c *ArticleController) ArticlesGyCategory() {

	category := c.GetString(":category")

	if len(category) > 0 {
		c.Data["active_category"] = category
	}
	userId := c.UserId()

	pagination := c.NewPagination()

	models.ArticlesGyCategory(userId, category, pagination)
	c.Data["pagination"] = pagination

	c.TplName = "index.html"
}

func (c *ArticleController) ArticleDetail() {
	id, _ := c.GetInt64(":id")
	ip := c.Ip()
	userId := c.UserId()
	beego.Error("id", id)
	c.TplName = "article_detail.html"

	if id <= 0 {
		c.StringError("文章不存在")
		return
	}

	article, error := models.ArticleById(id)

	if userId > 0 {
		hasLike, err := models.HasLikeArticle(id, userId, db.NewDB())
		if nil == err {
			article.HasLike = hasLike
		}
	}

	if nil != error {
		c.StringError("文章不存在")
	}else {
		success, _ := models.IncrViewCount(id, userId, ip)
		if success {
			article.ViewCount = article.ViewCount + 1
		}
		c.Data["article"] = article
		c.SetKeywords(article.Categories + "," + article.Tags)
		var subLength = models.ParameterIntValue("seo-description-length")
		c.SetDescription(article.ShortContent(subLength)).SetTitle(article.Title)
	}
}


func (c *ArticleController) EditArticle() {
	id, _ := c.GetInt64(":id")
	userId := c.UserId()
	c.TplName = "article_edit.html"

	if id <= 0 {
		c.StringError("文章不存在")
		return
	}

	if (userId == 0) {
		c.Redirect("/login", 302)
		return
	}

	article, error := models.ArticleById(id)

	if nil != error {
		c.StringError("文章不存在")
	}else {
		c.Data["article"] = article
		c.SetKeywords(article.Categories + "," + article.Tags)
		var subLength = models.ParameterIntValue("seo-description-length")
		c.SetDescription(article.ShortContent(subLength)).SetTitle(article.Title)
	}
}



func (c *ArticleController) UpdateArticle() {
	user := c.GetSession("user")
	if (nil == user) {
		c.Redirect("/login", 200)
	}

	id, err := c.GetInt64("id")

	var article entities.Article
	if nil == err {
		userId := c.UserId()
		content := c.GetString("content")
		title := c.GetString("title")
		tag := c.GetString("tag")
		categories := c.GetString("category")
		beego.Error("tag", tag, "categories", categories)
		article = entities.Article{Id:id, UserId:userId, Title:title, Content:content, CreatedAt:time.Now()}
		prt := &article
		prt.SetCategories(categories)
		prt.SetTags(tag)
		beego.Error("tag", article.Tags, "categories", article.Categories)
		exception, valid := models.ValidArticle(article)
		if valid {
			err = models.UpdateArticle(&article)
		}else {
			err = exception
		}

	}

	if (nil == err) {
		c.Redirect("../", 302)
	}else {
		beego.Error(err)
		c.StringError(err.Error())
		c.Data["article"] = article
		c.TplName = "article_create.html"
	}
}


func (c *ArticleController) New() {
	user := c.GetSession("user")

	if (nil == user) {
		c.Redirect("/login", 302)
	}else {
		c.TplName = "article_create.html"
	}
}

func (c *ArticleController) CreateArticle() {

	user := c.GetSession("user")
	if (nil == user) {
		c.Redirect("/login", 200)
	}

	userId := c.UserId()
	content := c.GetString("content")
	title := c.GetString("title")
	tag := c.GetString("tag")
	categories := c.GetString("category")

	beego.Debug("title", title, "categories", categories, "tag", tag, "content", content)

	article := entities.Article{UserId:userId, Title:title, Content:content, CreatedAt:time.Now()}
	err, valid := models.ValidArticle(article)
	if valid {
		prt := &article
		prt.SetCategories(categories)
		prt.SetTags(tag)
		err = models.SaveArticle(&article)
	}

	if (nil == err) {
		c.Redirect("../", 302)
	}else {
		beego.Error(err)
		c.StringError(err.Error())
		c.Data["article"] = article
		c.TplName = "article_create.html"
	}

}


func (c *ArticleController) Like() {


	articleId, err := c.GetInt64(":id")
	if nil != err {
		c.ErrorCodeJsonError(exception.NOT_EXIST_ARTICLE_ERROR)
		return
	}
	userId := c.UserId()

	if userId < 1 {
		c.ErrorCodeJsonError(exception.NOT_LOGIN)
		return
	}

	var incrCount int
	incrCount, err = models.IncrLikeCount(articleId, userId)

	if nil == err {
		c.JsonSuccess(incrCount)
	}else {
		c.JsonError("")
	}
}

