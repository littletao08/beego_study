package controllers
import (
"github.com/astaxie/beego"
"beego_study/models"
)
// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

type ResponseBody struct {
	Success bool
	Message string
	Code int
	Data interface{}
}

func (c *BaseController) Prepare() {
	categories,_ := models.Categories()
	c.Data["categories"] = categories
	c.Data["showRightBar"] = true
	response := ResponseBody{Success:true}
	c.Data["response"]=response
}
