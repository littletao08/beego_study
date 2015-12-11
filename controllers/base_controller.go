package controllers
import (
	"github.com/astaxie/beego"
	"beego_study/models"
	"github.com/gogather/com/log"
	"beego_study/db"
	"beego_study/entities"
	"net/url"
	"strings"
	"encoding/json"
	"beego_study/exception"
)
// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

type ResponseBody struct {
	Success bool
	Message interface{}
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
	userAgent := c.Ctx.Request.UserAgent()
	userAgent = strings.ToLower(userAgent)
	//是否是手机访问
	c.Data["isMobile"] = false
	if (strings.Contains(userAgent, "android") || strings.Contains(userAgent, "iphone")) {
		c.Data["isMobile"] = true
	}

	log.Bluef(args)
}


func (c *BaseController) NewPagination() *db.Pagination {
	page, err := c.GetInt("page")
	if nil != err {
		page = 1
	}
	pagination := db.NewPagination(page, 0, false)
	link, _ := url.ParseRequestURI(c.Ctx.Request.URL.String())
	pagination.SetUrl(link)
	return pagination
}


func (c *BaseController) CurrentUser() *entities.User {
	user := c.GetSession("user")
	if nil == user {
		return nil
	}
	var u, ok = user.(entities.User)
	if !ok {
		return nil
	}
	return &u
}

func (c *BaseController) UserId() int64 {
	user := c.CurrentUser()
	if nil == user {
		return 0
	}
	return user.Id
}
func (c *BaseController) StringError(message string) {
	response := new(ResponseBody)
	response.Code = -1
	response.Success = false
	response.Message = message
	result, err := json.Marshal(response)
	if nil == err {
		c.Data["message"] = string(result)
	}
}

func (c *BaseController) StringSuccess(message string) {
	response := new(ResponseBody)
	response.Code = 0
	response.Success = true
	response.Message = message
	result, err := json.Marshal(response)
	if nil == err {
		c.Data["message"] = string(result)
	}
}

func (c *BaseController) ErrorCodeJsonError(exception exception.ErrorCode) {
	response := new(ResponseBody)
	response.Code = exception.Code()
	response.Success = false
	response.Message = exception.Message()
	c.Data["json"] = response
	c.ServeJson()
}

func (c *BaseController) JsonError(message interface{}) {
	response := new(ResponseBody)
	response.Code = -1
	response.Success = false
	response.Message = message
	c.Data["json"] = response
	c.ServeJson()
}

func (c *BaseController) JsonSuccess(message interface{}) {
	response := new(ResponseBody)
	response.Code = 0
	response.Success = true
	response.Message = message
	c.Data["json"] = response
	c.ServeJson()
}


func (c *BaseController) Ip() string {
	return c.Ctx.Request.Header.Get("X-Real-Ip")
}