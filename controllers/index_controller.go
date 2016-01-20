package controllers
import (
	"beego_study/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	pagination := c.NewPagination()
	userId := c.CurrentUserId()
	models.AllArticles(userId, pagination)
	c.Data["pagination"] = pagination
	c.Data["user"]=c.CurrentUser()
	c.TplName = "index.html"

}
