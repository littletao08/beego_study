package controllers
import (
	"beego_study/models"
"github.com/astaxie/beego"
	"beego_study/entities"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	pagination := c.NewPagination()
	userId := c.UserId()
	models.AllArticles(userId, pagination)

	c.Data["pagination"] = pagination
	c.TplNames = "index.html"
	for _, v := range pagination.Data {
		article := v.(entities.Article)
		beego.Error("userid:",userId,"articleid:", article.Id,"hasLike", article.HasLike)
	}

}
