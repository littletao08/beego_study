package controllers
import (
	"beego_study/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	pagination := c.NewPagination()
	userId := c.UserId()
	models.AllArticles(userId, pagination)
	c.Data["pagination"] = pagination
	c.Data["user"]=c.CurrentUser()
	beego.Error("user:",c.Data["user"])
	c.TplName = "index.html"

}
