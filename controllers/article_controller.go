package controllers

import (
	"beego_study/models"
	"beego_study/entities"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Articles() {
	page, _ := c.GetInt(":page")
	articles, _ := models.Articles(page)
	c.Data["articles"] = articles
	c.TplNames = "index.html"
}

func (c *ArticleController) New() {
	c.TplNames = "article_create.html"
}

func (c *ArticleController) CreateArticle() {
	content := c.GetString("content")
	title := c.GetString("title")
	article := entities.Article{UserId:1, Title:title, Content:content}
	err := models.Save(&article)
	if (nil == err) {
		c.Data["json"] = true
	}else {
		c.Data["json"] = false

	}
	c.TplNames = "index.html"
}
