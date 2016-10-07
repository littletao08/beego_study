package controllers

import (
	"github.com/astaxie/beego"
	"beego_study/services"
	"beego_study/db"
	"beego_study/entities"
	"net/url"
	"strings"
	"encoding/json"
	"beego_study/exception"
	"io"
	"bytes"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"reflect"
	"beego_study/utils"
	"time"
	"beego_study/utils/redis"
	"beego_study/models"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

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

	c.Data["showLeftBar"] = true
	var keywords, _ = services.ParameterValue("index-keywords")
	c.Data["keywords"] = keywords
	var description, _ = services.ParameterValue("index-description")
	c.Data["description"] = description

	response := ResponseBody{Success:true}
	c.Data["response"] = response

	var args interface{}
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
	user := c.CurrentUser()
	if nil != user {
		c.Data["user"] = user
	}

	beego.Info("request-params:", args)
}

func (c *BaseController) Finish() {

}

func (c *BaseController) Render() error {
	if !c.EnableRender {
		return nil
	}
	rb, err := c.RenderBytes()

	if err != nil {
		return err
	} else {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		if err != nil {
			return err
		}
		miniRb, err := mini(rb)
		if err != nil {
			return err
		}
		c.Ctx.Output.Body(miniRb)
	}

	return nil
}

func mini(renderBytes []byte) ([]byte, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", func(_ *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
		_, err := io.Copy(w, r)
		return err
	})
	m.AddFunc("text/javascript", func(_ *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
		_, err := io.Copy(w, r)
		return err
	})
	r := bytes.NewBuffer(renderBytes)
	w := &bytes.Buffer{}
	err := html.Minify(m, w, r, nil)

	return w.Bytes(), err
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

func (c *BaseController) CurrentOpenUser() *entities.OpenUser {
	openUser := c.GetSession("openUser")
	if nil == openUser {
		return nil
	}
	var u, ok = openUser.(*entities.OpenUser)
	if !ok {
		return nil
	}
	return u
}

func (c *BaseController) SetCurrSession(sessionKey string, value interface{}) {
	c.SetSession(sessionKey, value)
}

func (c *BaseController) CurrentUserId() int64 {
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
	response.Message = exception.Error()
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *BaseController) JsonError(message interface{}) {
	response := new(ResponseBody)
	response.Code = -1
	response.Success = false
	response.Message = message
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *BaseController) JsonSuccess(message interface{}) {
	response := new(ResponseBody)
	response.Code = 0
	response.Success = true
	response.Message = message
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *BaseController) toJsonResponse(key string, value interface{}, extraKeyValues ... interface{}) {

	keyValues := []interface{}{}
	keyValues = append(keyValues, key, value)

	for _, v := range extraKeyValues {
		val := reflect.ValueOf(v)
		sInd := reflect.Indirect(val)
		if (sInd.Kind() == reflect.Slice) {
			var s = utils.ToSlice(v)
			keyValues = append(keyValues, s...)
		} else {
			keyValues = append(keyValues, v)
		}
	}

	/*if (len(keyValues) % 2 != 0) {

	}*/
	var data = make(map[string]interface{})
	len := len(keyValues)
	for i := 0; i < len; i += 2 {
		var k string
		value, ok := keyValues[i].(string);
		if ok {
			k = value
		}
		v := keyValues[i + 1]
		v = models.Rewrite(v);
		data[k] = v
	}

	c.Data["json"] = data
	c.ServeJSON()

}

func (c *BaseController) FailResponse(message string) {
	c.toJsonResponse("success", false, "message", message);
}

func (c *BaseController) SuccessResponse(extraKeyValues ... interface{}) {
	c.toJsonResponse("success", true, extraKeyValues);
}

func (c *BaseController) Ip() string {
	return c.Ctx.Request.Header.Get("X-Real-Ip")
}

func (c *BaseController) SetKeywords(keywords string) *BaseController {
	c.Data["keywords"] = keywords
	return c
}

func (c *BaseController) SetDescription(description string) *BaseController {
	c.Data["description"] = description
	return c
}

func (c *BaseController) SetTitle(title string) *BaseController {
	c.Data["title"] = title
	return c
}

func (c *BaseController) VerifyCaptcha() bool {
	captcha := c.GetString("captcha")
	captchaId := c.GetString("captcha_id")
	return cpt.Verify(captchaId, captcha)
}

func (c *BaseController) RecordLoginFailTimes() {
	sessionId := c.CruSession.SessionID()

	failTimes := services.ParameterIntValue("no_captcha_login_fail_times")

	failTimesCache := redis_util.IncrByWithTimeOut(sessionId, 1, int64(time.Second * 3))

	if failTimesCache >= failTimes {
		c.Data["showCaptcha"] = true
	}
}

func (c *BaseController) NeedCheckCaptcha() bool {
	sessionId := c.CruSession.SessionID()

	failTimes := services.ParameterIntValue("no_captcha_login_fail_times")

	failTimesCache := redis_util.IncrByWithTimeOut(sessionId, failTimes, int64(time.Second * 3))

	return failTimesCache > failTimes;
}

func (c *BaseController) SetCategories(userId int64) {
	categories, _ := services.UserCategories(userId)
	c.Data["categories"] = categories
}

func (c *BaseController) Host() string {

	/*var url = c.Ctx.Request.Referer();
	url = strings.Split(url, "//")[1]
	host := strings.Split(url, "/")[0]*/

	host := "www.threeperson.com"
	return host

}

