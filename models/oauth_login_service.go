package models

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"beego_study/utils"
	"net/http"
	"encoding/json"
	"beego_study/entities"
	"time"
	"errors"
)

var AuthConfig config.Configer

func init() {

	var appPath = beego.AppConfigPath

	appPath = beego.Substr(appPath, 0, strings.LastIndex(appPath, "/") + 1)
	beego.Debug("appPath:", appPath)
	config, err := config.NewConfig(beego.AppConfigProvider, appPath + "/auth_login.conf")
	if err != nil {
		beego.Error(config)
		panic("auth_login.conf load fail !")
	}
	AuthConfig = config

}

func QueryToken(authCode string) (map[string]string, error) {
	params := make(map[string]string)
	params["grant_type"] = "authorization_code"
	params["client_id"] = AuthConfig.String("app_id")
	params["client_secret"] = AuthConfig.String("app_key")
	params["code"] = authCode
	params["redirect_uri"] = AuthConfig.String("token_redirect_uri")

	var baseUrl = AuthConfig.String("access_token_url")

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

func QueryOpenId(accessToken string) (map[string]string, error) {

	var baseUrl = AuthConfig.String("openid_url")
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

	beego.Debug("******************queryOpenID:before", content)
	start := strings.Index(content, "{") - 1
	end := strings.LastIndex(content, "}") + 1
	content = beego.Substr(content, start, end - start)
	beego.Debug("******************queryOpenID:after", content)

	var paramMap map[string]string

	err = json.Unmarshal([]byte(content), &paramMap)
	return paramMap, err

}

func OpenUserInfo(accessToken string, openId string) (*entities.OpenUser, error) {
	var baseUrl = AuthConfig.String("get_user_info_url")
	params := make(map[string]string)
	params["access_token"] = accessToken
	params["openid"] = openId
	params["oauth_consumer_key"] = AuthConfig.String("app_id")

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

	beego.Debug("******************OpenUserInfo", content)

	var paramMap map[string]interface{}

	err = json.Unmarshal([]byte(content), &paramMap)

	if len(paramMap) == 0 {
		return nil, errors.New("open_user_info get fail")
	}

	age, ok := paramMap["age"].(int)
	if !ok {
		age = 0
	}
	city, _ := paramMap["city"].(string)
	province, _ := paramMap["province"].(string)
	nick, _ := paramMap["nickname"].(string)
	gender, _ := paramMap["gender"].(string)
	head, _ := paramMap["figureurl_qq_1"].(string)
	year, ok := paramMap["year"].(int)
	if !ok {
		year = 0
	}

	sex := 1
	if "女" == gender {
		sex = 2
	}

	if year > 0 {
		currYear := time.Now().Year()
		age = currYear - year
	}

	openUser := new(entities.OpenUser)

	openUser.OpenId = openId
	openUser.Nick = nick
	openUser.Age = age
	openUser.City = city
	openUser.Province = province
	openUser.Sex = sex
	openUser.Type = entities.OPEN_USER_TYPE_QQ
	openUser.Head = head

	return openUser, err
}


