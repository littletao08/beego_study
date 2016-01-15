package controllers

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"beego_study/utils"
"io/ioutil"
	"encoding/json"
)

type AuthLoginController struct {
	BaseController
}

var authConfig config.ConfigContainer

func init() {

	var appPath = beego.AppConfigPath

	appPath = beego.Substr(appPath, 0, strings.LastIndex(appPath, "/") + 1)

	config, err := config.NewConfig(beego.AppConfigProvider, appPath + "/auth_login.conf")
	if err != nil {
		beego.Error(config)
		panic("auth_login.conf load fail !")
	}
	authConfig = config

}
func (c *AuthLoginController) QQAuth() {

	params := make(map[string]string)
	params["client_id"] = authConfig.String("app_id")
	params["redirect_uri"] = authConfig.String("auth_redirect_uri")
	params["response_type"] = "code"
	var state = "111";
	params["state"] = state
	params["scope"] = authConfig.String("scope")

	var baseUrl = authConfig.String("authorize_url")

	subUrl := utils.BuildQueryString(params)

	var fullRequestStr = baseUrl + subUrl
	c.Redirect(fullRequestStr, 302)
}

func (c *AuthLoginController) QQToken() {

	code := c.GetString("code")

	//获取token
	tokenRes, err := queryToken(code)
	beego.Debug("****************tokenRes:",tokenRes,"****************")
	if (nil != err ) {
		c.Redirect("login.html", 302)
		return
	}

	accessToken := tokenRes["access_token"]
	beego.Debug("****************accessToken:",accessToken,"****************")
	if len(accessToken) <= 0 {
		c.Redirect("login.html", 302)
		return
	}

	//获取openid
	openIdRes,err := queryOpenID(accessToken)
	beego.Debug("****************openIdRes:",openIdRes,"****************")
	if (nil != err ) {
		c.Redirect("login.html", 302)
		return
	}

	openId := openIdRes["openid"]
	if len(openId) <= 0 {
		c.Redirect("login.html", 302)
		return
	}

	//获取user_info
	userInfoRes,err := OpenUserInfo(accessToken, openId)
	beego.Debug("****************userInfoRes:",userInfoRes,"****************")

	if (nil != err ) {
		c.Redirect("login.html", 302)
		return
	}

	beego.Error("****************",userInfoRes,"****************")


	c.TplNames = "index.html"
}

func queryToken(authCode string) (map[string]string, error) {
	params := make(map[string]string)
	params["grant_type"] = "authorization_code"
	params["client_id"] = authConfig.String("app_id")
	params["client_secret"] = authConfig.String("app_key")
	params["code"] = authCode
	params["redirect_uri"] = authConfig.String("token_redirect_uri")

	var baseUrl = authConfig.String("access_token_url")

	subUrl := utils.BuildQueryString(params)

	var fullRequestStr = baseUrl + subUrl

	resp, err := http.Get(fullRequestStr)
	if nil != err {
		beego.Error(err)
		return nil, err
	}

	paramMap, err := utils.ExtractResponse(resp.Body)

	defer resp.Body.Close()

	return paramMap, err
}

func queryOpenID(accessToken string) (map[string]string, error) {

	var baseUrl = authConfig.String("openid_url")
	var fullRequestStr = baseUrl + "?access_token=" + accessToken

	resp, err := http.Get(fullRequestStr)
	if nil != err {
		beego.Error(err)
		return nil, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if nil != err {
		beego.Error(err)
		return nil, err;
	}

	content := string(result)
	if len(content) == 0 {
		beego.Error(err)
		return nil, err;
	}

	beego.Debug("******************queryOpenID:before",content)
	content = beego.Substr(content,strings.Index(content,"{")-1,strings.LastIndex(content,"}")+1)
	beego.Debug("******************queryOpenID:after",content)

	var paramMap map[string]string

	err = json.Unmarshal([]byte(content),&paramMap)
	return paramMap, err

}


func OpenUserInfo(accessToken string,openId string) (map[string]string,error) {
	var baseUrl = authConfig.String("get_user_info_url")
	params := make(map[string]string)
	params["access_token"] = accessToken
	params["openid"] = openId
	params["oauth_consumer_key"] = authConfig.String("app_id")

	subUrl := utils.BuildQueryString(params)

	var fullRequestStr = baseUrl + subUrl

	resp, err := http.Get(fullRequestStr)
	if nil != err {
		beego.Error(err)
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if nil != err {
		beego.Error(err)
		return nil, err;
	}

	content := string(result)
	if len(content) == 0 {
		beego.Error(err)
		return nil, err;
	}

	beego.Debug("******************OpenUserInfo",content)

	var paramMap map[string]string

	err = json.Unmarshal([]byte(content),&paramMap)

	return paramMap, err

}


