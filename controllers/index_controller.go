package controllers
import (
	"beego_study/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	pagination := c.NewPagination()
	userId := c.UserId()
	models.AllArticles(userId, pagination)
	c.Data["pagination"] = pagination
	c.TplName = "index.html"

}
