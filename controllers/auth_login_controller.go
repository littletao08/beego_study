package controllers

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
	"bytes"
)

type AuthLoginController struct {
	BaseController
}

var authConfig config.ConfigContainer

func init() {

	beego.Debug(beego.AppPath)

	config, err := config.NewConfig(beego.AppConfigProvider, beego.AppPath + "/conf/auth_login.conf")
	if err != nil {
		beego.Error(config)
		panic("auth_login.conf load fail !")
	}
	authConfig = config

}
func (c *AuthLoginController) QQAuth() {

	params := make(map[string]string)
	params["client_id"] = authConfig.String("app_id")
	params["redirect_uri"] = authConfig.String("redirect_uri")
	params["response_type"] = "code"
	var state = "111";
	params["state"] = state
	params["scope"] = authConfig.String("scope")

	var baseUrl = authConfig.String("authorize_url")

	reqStr := bytes.NewBufferString("?")

	for key, val := range params {
		reqStr.WriteString(key)
		reqStr.WriteString("=")
		reqStr.WriteString(val)
		reqStr.WriteString("&")
	}
	var fullRequestStr = baseUrl + reqStr.String()
	c.Redirect(fullRequestStr, 302)
}

func (c *AuthLoginController) QQToken() {

	code := c.GetString("code")
	params := make(map[string]string)
	params["grant_type"] = code
	params["client_id"] = authConfig.String("app_id")
	params["client_secret"] = authConfig.String("app_key")
	params["state"] = code
	params["redirect_uri"] = authConfig.String("redirect_uri")

	var baseUrl = authConfig.String("access_token_url")

	reqStr := bytes.NewBufferString("?")

	for key, val := range params {
		reqStr.WriteString(key)
		reqStr.WriteString("=")
		reqStr.WriteString(val)
		reqStr.WriteString("&")
	}
	var fullRequestStr = baseUrl + reqStr.String()
	c.Redirect(fullRequestStr, 302)
}



