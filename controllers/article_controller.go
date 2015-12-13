package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"github.com/gogather/com/log"
	"time"
	"github.com/astaxie/beego"
	"beego_study/exception"
	"beego_study/db"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Articles() {
	page, _ := c.GetInt("page")
	articles, _ := models.Articles(page)
	c.Data["articles"] = articles
	c.TplNames = "index.html"
}

func (c *ArticleController) ArticleDetail() {
	id, _ := c.GetInt64(":id")
	ip := c.Ip()
	userId := c.UserId()
	beego.Error("id", id)
	c.TplNames = "article_detail.html"

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
	}
}

func (c *ArticleController) New() {
	user := c.GetSession("user")
	log.Redln("*******************", user)
	if (nil == user) {
		c.Redirect("/login", 302)
	}else {
		c.TplNames = "article_create.html"
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

	article := entities.Article{UserId:userId, Title:title, Categories: categories, Tags:tag, Content:content, CreatedAt:time.Now()}
	err := models.SaveArticle(&article)
	if (nil == err) {
		c.Redirect("../", 302)
	}else {
		beego.Error("文章创建失败")
		c.StringError("文章创建失败")
		c.Data["article"] = article
		c.TplNames = "article_create.html"
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

