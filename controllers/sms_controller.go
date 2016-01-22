package controllers
import (
	"encoding/json"
	"beego_study/models"
	"beego_study/utils/redis"
	"strings"
	"time"
	"log"
)
var smsCodeLength, smsCodeVerifyCount int
var smsCodePrefix, smsCodeVerifyCountPrefix string
var smsCodeTimeOut int64

func init() {
	smsCodeLength = 6
	smsCodePrefix = "sms_code_"
	//限制同一个手机号同一天只能发送的次数
	smsCodeVerifyCount = models.AuthConfig.DefaultInt("sms-code-verify-count", 10)
	smsCodeVerifyCountPrefix = "sms_code_verify_count_"
	//验证码有效期
	smsCodeTimeOut = models.AuthConfig.DefaultInt64("sms-code-time-out", 600)
}
type SmsController struct {
	BaseController
}

func (sms *SmsController) Send() {
	mobile := sms.GetString("mob")
	smsResponse := &models.SmsResponse{}
	if verify, ct := verifyMaxCount(mobile); verify {
		//生成验证码,默认是六位
		smsCode := models.RandomSmsCode(smsCodeLength)

		smsRequest := &models.SmsRequest{MobilePhoneNumber:mobile, Content:smsCode}
		smsResponse = models.SendRegisterSms(smsRequest)
		bytes, _ := json.Marshal(smsResponse)
		jsonValue := string(bytes)
		//发送并且响应成功以后
		if smsResponse != nil && smsResponse.Code == 0 {
			redis_util.Set(smsCodePrefix + mobile, smsCode, smsCodeTimeOut)
			//设置每天短信验证码的次数
			ct = ct + 1
			timeStr := time.Now().Format("2006-01-02")
			log.Printf("短信验证码发送成功,当天发送次数:%d", ct)
			redis_util.Set(smsCodeVerifyCountPrefix + timeStr + mobile, ct, 86400)
		}
		sms.Data["json"] = jsonValue
	}else {
		smsResponse.ResMessage = "发送验证码次数已经超过上限"
		bytes, _ := json.Marshal(smsResponse)
		sms.Data["json"] = string(bytes)
	}
	sms.ServeJSON()
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

