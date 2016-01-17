package controllers

import (
	"github.com/astaxie/beego"
	"beego_study/utils"
	"beego_study/models"
	"beego_study/entities"
)

type AuthLoginController struct {
	BaseController
}


func (c *AuthLoginController) QqAuth() {

	params := make(map[string]string)
	params["client_id"] = models.AuthConfig.String("app_id")
	params["redirect_uri"] = models.AuthConfig.String("auth_redirect_uri")
	params["response_type"] = "code"
	var state = "111";
	params["state"] = state
	params["scope"] = models.AuthConfig.String("scope")

	var baseUrl = models.AuthConfig.String("authorize_url")

	subUrl := utils.BuildQueryString(params)

	var fullRequestStr = baseUrl + subUrl
	c.Redirect(fullRequestStr, 302)
}

func (c *AuthLoginController) QqToken() {

	code := c.GetString("code")

	//获取token
	tokenRes, err := models.QueryToken(code)
	beego.Debug("****************tokenRes:", tokenRes, "****************")
	if (nil != err ) {
		c.Redirect("/login", 302)
		return
	}

	accessToken := tokenRes["access_token"]
	beego.Debug("****************accessToken:", accessToken, "****************")
	if len(accessToken) <= 0 {
		c.Redirect("/login", 302)
		return
	}

	//获取openid
	openIdRes, err := models.QueryOpenId(accessToken)
	beego.Debug("****************openIdRes:", openIdRes, "****************")
	if (nil != err ) {
		c.Redirect("login.html", 302)
		return
	}

	openId := openIdRes["openid"]
	if len(openId) <= 0 {
		c.Redirect("/login", 302)
		return
	}

	//获取user_info
	openUser, err := models.OpenUserInfo(accessToken, openId)
	beego.Debug("****************userInfoRes:", openUser, "****************")

	if (nil != err || len(openUser) == 0 ) {
		c.Redirect("/login", 302)
		return
	}

	oauth := new(entities.OpenUser)
	oauth.OpenId = openId
	oauth.Age =

	beego.Error("****************", openUser, "****************")

	c.Redirect("/", 302)
}



