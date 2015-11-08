package controllers
import "beego_study/models"

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	categories,_ := models.Categories()
	c.Data["categories"] = categories
	c.TplNames = "index.html"
}
