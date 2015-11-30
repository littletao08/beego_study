package controllers
import (
	"beego_study/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	pagination := c.NewPagination()
	userId := c.GetUserId()
	models.AllArticles(userId,pagination)
	c.Data["pagination"] = pagination
	c.TplNames = "index.html"
}
