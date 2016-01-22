/**
短信相关的服务,基础的服务,包括生成短信验证码
调用渠道接口发送短信内容.
 */
package models
import (
	"github.com/astaxie/beego"
	"encoding/json"
	"strings"
	"strconv"
	"math/rand"
	"time"
	"net/http"
	"log"
	"go-exercise/go/src/fmt"
)

type SmsResponse struct {
	SmsId      int32  `json:"smsId"`
	ResMessage string `json:"error"`
	Code       int32  `json:"code"`
	Success    bool `json:"success"`
}

type SmsRequest struct {
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
	Content           string `json:"content"`
}
var registerTemplate, appId, appKey, contentType, smsUrl string

func init() {
	registerTemplate = AuthConfig.String("sms-register-template")
	contentType = AuthConfig.String("sms-Content-Type")
	appKey = AuthConfig.String("sms-X-Bmob-REST-API-Key")
	appId = AuthConfig.String("sms-X-Bmob-Application-Id")
	smsUrl = AuthConfig.String("sms-bmob-url")
}

func SendRegisterSms(request *SmsRequest) *SmsResponse {
	request.Content = fmt.Sprintf(registerTemplate, request.Content)
	return Send(request)
}

//生成一个纯数字的验证码.
func RandomSmsCode(length int) string {
	values := make([]string, length)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range values {
		value := rand.Intn(10)
		values[i] = strconv.Itoa(value)
	}
	return strings.Join(values, "")
}

//发送短信内容接口
func Send(request *SmsRequest) (*SmsResponse) {
	client := &http.Client{}

	jsonValues, err := json.Marshal(request)
	if err != nil {
		return nil
	}
	params := string(jsonValues)
	log.Printf("params : %s \n", params)

	req, err := http.NewRequest("POST", smsUrl, strings.NewReader(params))
	if err != nil {
		beego.Error("短信请求发送失败")
		return nil
	}

	//设置Head参数
	req.Header.Add("X-Bmob-Application-Id", appId)
	req.Header.Add("X-Bmob-REST-API-Key", appKey)
	req.Header.Add("Content-Type", contentType)
	log.Printf("sms send request : %+v \n", req)
	resp, err := client.Do(req)
	if err != nil {
		beego.Error("短信请求发送失败")
		return nil
	}
	defer resp.Body.Close()
	smsResponse := &SmsResponse{}
	err = json.NewDecoder(resp.Body).Decode(smsResponse)
	if err != nil {
		beego.Error("短信内容解析失败")
		return nil
	}
	log.Printf("短信验证码发送结果:%v+ \n", smsResponse)
	return smsResponse
}
