package controllers

import (
	"beego_study/models"
	"beego_study/entities"
	"github.com/gogather/com/log"
	"time"
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
	//id, _ := c.GetInt(":id")
	//c.Data["showRightBar"] = false
	c.TplNames = "article_detail.html"
}

func (c *ArticleController) New() {
	user := c.GetSession("user")
	log.Redln("*******************",user)
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
	article := entities.Article{UserId:1, Title:title, Content:content,CreatedAt:time.Now()}
	err := models.Save(&article)
	if (nil == err) {
		c.Data["json"] = true
	}else {
		c.Data["json"] = false
	}
	c.ServeJson()
}
