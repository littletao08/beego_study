package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"github.com/gogather/com/log"
	"time"
	"github.com/astaxie/beego"
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
	beego.Error("id", id)
	c.TplNames = "article_detail.html"
	if id <= 0 {
		c.SetError("文章不存在")
		return
	}

	article, error := models.ArticleById(id)
	if nil != error {
		c.SetError("文章不存在")
	}else {
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

	content := c.GetString("content")
	title := c.GetString("title")
	tag := c.GetString("tag")
	category := c.GetString("category")

	userId := c.GetUserId()

	beego.Error("title",title,"category",category,"tag",tag,"content",content)

	article := entities.Article{UserId:userId, Title:title,Tag:tag, Content:content, CreatedAt:time.Now()}
	err := models.Save(&article)
	if (nil == err) {
		c.Redirect("/",302)
	}else {
		c.SetError("文章创建失败")
		c.TplNames = "article_create.html"
	}
}
