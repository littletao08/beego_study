package controllers
import (
	"github.com/astaxie/beego"
	"beego_study/models"
	"github.com/gogather/com/log"
	"github.com/astaxie/beego/utils/pagination"
)
// Controller基类继承封装
type BaseController struct {
	beego.Controller
	pagination.Paginator
}

type ResponseBody struct {
	Success bool
	Message string
	Code    int
	Data    interface{}
}

func (c *BaseController) Prepare() {
	categories, _ := models.Categories()
	c.Data["categories"] = categories
	c.Data["showRightBar"] = true
	response := ResponseBody{Success:true}
	c.Data["response"] = response
	var args string
	method := c.Ctx.Request.Method
	if ("GET" == method) {
		args = c.Ctx.Request.RequestURI
	}else {
		args = c.Ctx.Request.Form.Encode();
	}
	log.Bluef(args)
}
