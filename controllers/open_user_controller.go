package controllers

import (
	"github.com/astaxie/beego"
	"beego_study/utils"
	"beego_study/services"
	"errors"
	"beego_study/entities"
)

type OpenUserController struct {
	BaseController
}

func (c *OpenUserController) QqAuth() {
	params := make(map[string]string)
	params["client_id"] = services.AuthConfig.String("app_id")
	params["redirect_uri"] = services.AuthConfig.String("auth_redirect_uri")
	params["response_type"] = "code"
	var state = "111";
	params["state"] = state
	params["scope"] = services.AuthConfig.String("scope")

	var baseUrl = services.AuthConfig.String("authorize_url")

	subUrl := utils.BuildQueryString(params)

	var fullRequestStr = baseUrl + subUrl
	c.Redirect(fullRequestStr, 302)
}

func (c *OpenUserController) NewOrBindUser()  {
	c.Data["showLeftBar"] = false
	c.TplName="register.html"
}


func (c *OpenUserController) QqToken() {

	code := c.GetString("code")

	var loginPageUrl = "/users/login"
	//获取token
	tokenRes, err := services.QueryToken(code)
	beego.Debug("****************tokenRes:", tokenRes, "****************")
	if (nil != err ) {
		c.Redirect(loginPageUrl, 302)
		return
	}

	accessToken := tokenRes["access_token"]
	beego.Debug("****************accessToken:", accessToken, "****************")
	if len(accessToken) <= 0 {
		c.Redirect(loginPageUrl, 302)
		return
	}

	//获取openid
	openIdRes, err := services.QueryOpenId(accessToken)
	beego.Debug("****************openIdRes:", openIdRes, "****************")
	if (nil != err ) {
		beego.Error(err)
		c.Redirect("login.html", 302)
		return
	}

	openId := openIdRes["openid"]
	if len(openId) <= 0 {
		beego.Error(errors.New("openid["+openId+"] error"))
		c.Redirect(loginPageUrl, 302)
		return
	}

	//获取user_info
	openUser, err := services.OpenUserInfo(accessToken, openId)
	beego.Debug("err:",err,"userInfoRes:", openUser)

	if (nil != err ) {
		beego.Error(err)
		c.Redirect(loginPageUrl, 302)
		return
	}

	err = services.SaveOrUpdateOpenUser(openUser)
	if nil != err {
		beego.Error(err)
		c.Redirect(loginPageUrl, 302)
		return
	}

	openUser,_ = services.OpenUser(openId,entities.OPEN_USER_TYPE_QQ)
	beego.Debug("openUser:",openUser)
	//绑定了账号则跳转到首页,否则跳转到注册或绑卡页面
    if openUser.HasBindUser() {
		userId := openUser.UserId
		user,_ := services.User(userId)
		c.SetCurrSession("user", user)
	}else {
		c.SetSession("openUser",openUser)
		c.Redirect("/users/oauth_register", 302)
		return
	}

	c.Redirect("/", 302)
}



