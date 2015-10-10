package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.Data["name"] = "张利达"
	c.Data["email"] = "lida.zhang.cj@gmail.com"
	c.TplNames = "index.html"
}
