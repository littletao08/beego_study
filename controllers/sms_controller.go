package controllers
import (
	"encoding/json"
	"beego_study/models"
	"beego_study/utils/redis"
	"strings"
	"time"
)
var smsCodeLength, smsCodeVerifyCount int
var smsCodePrefix, smsCodeVerifyCountPrefix string

func init() {
	smsCodeLength = models.AuthConfig.DefaultInt("sms-code-length", 6)
	smsCodePrefix = models.AuthConfig.DefaultString("sms-code-prefix", "sms_code_")
	smsCodeVerifyCount = models.AuthConfig.DefaultInt("sms-code-verify-count", 10)
	smsCodeVerifyCountPrefix = models.AuthConfig.DefaultString("sms-code-verify-count-prefix", "sms_code_verify_count_")
}
type SmsController struct {
	BaseController
}

func (sms *SmsController) Send() string {
	mobile := sms.GetString("mob")
	smsResponse := &models.SmsResponse{}
	if verify, ct := verifyMaxCount(mobile); verify {
		//生成验证码,默认是六位
		smsCode := models.RandomSmsCode(smsCodeLength)

		smsRequest := &models.SmsRequest{MobilePhoneNumber:mobile, Content:smsCode}
		smsResponse = models.Send(smsRequest)
		bytes, _ := json.MarshalIndent(smsResponse, " ", "")
		jsonValue := string(bytes)
		//发送并且响应成功以后
		if smsResponse != nil && smsResponse.Code == 0 {
			redis_util.Set(smsCodePrefix + mobile, sms, 600)
			//设置每天短信验证码的次数
			ct = + 1
			timeStr := time.Now().Format("2006-01-02")
			redis_util.Set(smsCodeVerifyCountPrefix + timeStr + mobile, ct, 86400)

		}
		return jsonValue
	}else {
		smsResponse.ResMessage = "发送验证码次数已经超过上限"
		bytes, _ := json.MarshalIndent(smsResponse, " ", "")
		return string(bytes)
	}
}

func (sms *SmsController) VerifySmsCode() {
	mobile := sms.GetString("mob")
	reSmsCode := sms.GetString("code")
	var smsCode string
	if err := redis_util.Get(smsCodePrefix + mobile, &smsCode); err != nil {
		//验证码不正确
	}
	if strings.EqualFold(reSmsCode, smsCode) {
		//验证码验证通过了
	}else {
		//验证码不正确
	}
}

func verifyMaxCount(mobile string) (bool, int) {
	var verifyCount int
	timeStr := time.Now().Format("2006-01-02")
	if err := redis_util.Get(smsCodeVerifyCountPrefix + timeStr + mobile, &verifyCount); err != nil {
		return true, 0
	}
	return smsCodeVerifyCount > verifyCount, verifyCount
}

