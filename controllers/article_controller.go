package controllers

import (
	"beego_study/services"
	"beego_study/entities"
	"time"
	"github.com/astaxie/beego"
	"beego_study/exception"
	"beego_study/db"
	"strconv"
)

type ArticleController struct {
	BaseController
}
func (c *ArticleController) Articles() {
	pagination := c.NewPagination()
	userId := c.CurrentUserId()
	services.AllArticles(userId, pagination)
	c.Data["pagination"] = pagination
	c.TplName = "index.html"

}

func (c *ArticleController) ArticlesGyCategory() {

	category := c.GetString(":category")

	if len(category) > 0 {
		c.Data["active_category"] = category
	}
	userId := c.CurrentUserId()

	pagination := c.NewPagination()

	services.ArticlesGyCategory(userId, category, pagination,false)
	c.Data["pagination"] = pagination

	c.TplName = "index.html"
}


func (c *ArticleController) ArticlesGyUserIdAndCategory() {

	userIdStr := c.GetString(":userId")
	category := c.GetString(":category")
	userId,_ := strconv.ParseInt(userIdStr,10,64)
	if len(category) > 0 {
		c.Data["active_category"] = category
	}

	pagination := c.NewPagination()

	services.ArticlesGyCategory(userId, category, pagination,true)
	c.Data["pagination"] = pagination
	c.SetCategories(userId)

	c.TplName = "user_home.html"
}

func (c *ArticleController) ArticleDetail() {
	id, _ := c.GetInt64(":id")
	ip := c.Ip()
	userId,_ := c.GetInt64(":userId")

	c.TplName = "article_detail.html"

	if id <= 0 {
		c.StringError("文章不存在")
		c.Ctx.Redirect(302, "/")
		return
	}

	article, error := services.ArticleById(id)

	if userId > 0 {
		hasLike, err := services.HasLikeArticle(id, userId, db.NewDB())
		if nil == err {
			article.HasLike = hasLike
		}
	}

	if nil != error {
		c.StringError("文章不存在")
	}else {
		success, _ := services.IncrViewCount(id, userId, ip)
		if success {
			article.ViewCount = article.ViewCount + 1
		}
		c.Data["article"] = article

		c.SetCategories(article.UserId)

		c.SetKeywords(article.Categories + "," + article.Tags)
		var subLength = services.ParameterIntValue("seo-description-length")
		c.SetDescription(article.ShortContent(subLength)).SetTitle(article.Title)
	}
}


func (c *ArticleController) EditArticle() {
	id, _ := c.GetInt64(":id")
	userId := c.CurrentUserId()
	c.TplName = "article_edit.html"

	beego.Error("id:",id)
	if id <= 0 {
		c.StringError("文章不存在")
		return
	}

	if (userId <= 0) {
		c.Redirect("/users/login", 302)
		return
	}

	article, error := services.ArticleByIdAndUserId(id,userId)

	if nil != error {
		c.StringError("文章不存在")
	}else {
		c.Data["article"] = article
		c.SetKeywords(article.Categories + "," + article.Tags)
		var subLength = services.ParameterIntValue("seo-description-length")
		c.SetDescription(article.ShortContent(subLength)).SetTitle(article.Title)
	}
}



func (c *ArticleController) UpdateArticle() {

	user := c.GetSession("user")
	if (nil == user) {
		c.Redirect("/users/login", 200)
	}

	id, err := c.GetInt64("id")
	beego.Error("id:",id)

	var article entities.Article
	if nil == err {
		userId := c.CurrentUserId()
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
		exception, valid := services.ValidArticle(article)
		if valid {
			err = services.UpdateArticle(&article)
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
		c.Redirect("/users/login", 302)
	}else {
		c.TplName = "article_create.html"
	}
}

func (c *ArticleController) CreateArticle() {

	user := c.GetSession("user")
	if (nil == user) {
		c.Redirect("/users/login", 200)
	}

	userId := c.CurrentUserId()
	content := c.GetString("content")
	title := c.GetString("title")
	tag := c.GetString("tag")
	categories := c.GetString("category")

	beego.Debug("title", title, "categories", categories, "tag", tag, "content", content)

	article := entities.Article{UserId:userId, Title:title, Content:content, CreatedAt:time.Now()}
	err, valid := services.ValidArticle(article)
	if valid {
		prt := &article
		prt.SetCategories(categories)
		prt.SetTags(tag)
		err = services.SaveArticle(&article)
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
	userId := c.CurrentUserId()

	if userId < 1 {
		c.ErrorCodeJsonError(exception.NOT_LOGIN)
		return
	}

	var incrCount int
	incrCount, err = services.IncrLikeCount(articleId, userId)

	if nil == err {
		c.JsonSuccess(incrCount)
	}else {
		c.JsonError("")
	}
}

