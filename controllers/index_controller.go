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
	models.AllArticles(0, pagination)
	models.SetLikeSign(pagination,userId)
	categories, _ := models.UserCategories(0)
	c.Data["categories"] = categories
	c.Data["pagination"] = pagination
	c.Data["user"]=c.CurrentUser()
	c.TplName = "index.html"

}
