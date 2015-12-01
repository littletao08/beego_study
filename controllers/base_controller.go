package controllers
import (
	"github.com/astaxie/beego"
	"beego_study/models"
	"github.com/gogather/com/log"
	"beego_study/db"
	"beego_study/entities"
	"net/url"
	"regexp"
)
// Controller基类继承封装
type BaseController struct {
	beego.Controller
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
	userAgent := c.Ctx.Request.UserAgent()
	patten := ";\\s?(\\S*?\\s?\\S*?)\\s?(Build)?/"
	reg := regexp.MustCompile(patten)
	var matchData = reg.Find([]byte(userAgent))
	userAgent = string(matchData[:len(matchData)])
	//是否是手机访问
	c.Data["isMobile"] = false
	if (len(userAgent) > 0) {
		c.Data["isMobile"] = true
		beego.Debug("**********user-agent:%s************", userAgent)
	}

	log.Bluef(args)
}


func (c *BaseController) NewPagination() *db.Pagination {

	page, err := c.GetInt("page")
	if nil != err {
		page = 1
	}
	log.Redln("page", page)
	pagination := db.NewPagination(page, 0, false)
	link, _ := url.ParseRequestURI(c.Ctx.Request.URL.String())
	pagination.SetUrl(link)
	return pagination
}


func (c *BaseController) GetUser() *entities.User {
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

func (c *BaseController) GetUserId() int64 {
	user := c.GetUser()
	if nil == user {
		return -1
	}
	return user.Id
}