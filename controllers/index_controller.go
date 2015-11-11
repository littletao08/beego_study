package controllers
import "beego_study/models"

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.Data["lastArticle"],_ = models.LastArticle()
	c.TplNames = "index.html"
}
