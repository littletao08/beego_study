package controllers
import (
	"beego_study/services"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	pagination := c.NewPagination()
	userId := c.CurrentUserId()
	services.AllArticles(0, pagination)
	services.SetLikeSign(pagination,userId)
	categories, _ := services.UserCategories(0)
	//推荐博客
	likeBlogs := services.TopLikeUsers()
	c.Data["likeBlogs"] = likeBlogs
	//精华文章
	likeArticles := services.TopLikeArticles()
	c.Data["likeArticles"] = likeArticles


	c.Data["categories"] = categories
	c.Data["pagination"] = pagination
	c.Data["user"]=c.CurrentUser()
	c.TplName = "index.html"
}


func (c *IndexController) Images(){
	c.Ctx.Request.Header.Set("Referer","http://www.threeperson.com")
	c.Redirect("http://threeperson.oss-cn-shanghai.aliyuncs.com/111.png",302);
}
