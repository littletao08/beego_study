package controllers
import (
	"encoding/json"
	"beego_study/models"
	"beego_study/utils/redis"
	"strings"
	"time"
	"log"
	"beego_study/entities"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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

func (controller *SmsController) MobRegister() {
	controller.Data["showLeftBar"] = false
	controller.TplName = "mob_register.html"
}

func (sms *SmsController) Send() {
	mobile := sms.GetString("mobile")
	smsResponse := &models.SmsResponse{}
	valid := validation.Validation{}

	if !valid.Mobile(mobile, "mobile").Ok {
		smsResponse.ResMessage = "手机号格式不正确!"
		bytes, _ := json.Marshal(smsResponse)
		sms.Data["json"] = string(bytes)
		sms.ServeJSON()
		return
	}
	//上限检测
	if verify, ct := verifyMaxCount(mobile); verify {
		//生成验证码,默认是六位
		smsCode := models.RandomSmsCode(smsCodeLength)

		smsRequest := &models.SmsRequest{MobilePhoneNumber:mobile, Content:smsCode}
		smsResponse = models.SendRegisterSms(smsRequest)
		bytes, _ := json.Marshal(smsResponse)
		//发送并且响应成功以后
		if smsResponse != nil && smsResponse.Code == 0 {
			redis_util.Set(smsCodePrefix + mobile, smsCode, smsCodeTimeOut)
			//设置每天短信验证码的次数
			ct = ct + 1
			timeStr := time.Now().Format("2006-01-02")
			log.Printf("短信验证码发送成功,当天发送次数:%d", ct)
			redis_util.Set(smsCodeVerifyCountPrefix + timeStr + mobile, ct, 86400)
		}else {
			smsResponse.ResMessage = "验证码发送失败,请稍后重试"
		}
		jsonValue := string(bytes)
		sms.Data["json"] = jsonValue
	}else {
		smsResponse.ResMessage = "发送验证码次数已经超过上限"
		bytes, _ := json.Marshal(smsResponse)
		sms.Data["json"] = string(bytes)
	}
	sms.ServeJSON()
}

func (this *SmsController) VerifySmsCode() {
	mobile := this.GetString("mobile")
	reSmsCode := this.GetString("smsCode")
	pwd := this.GetString("password")
	repwd := this.GetString("repassword")
	//失败的时候免用户输入手机号
	this.Data["mobile"] = mobile

	beego.Debug("手机号注册参数:mobile:", mobile, "smsCode:", reSmsCode)
	//验证密码和重复密码是否一致
	if !strings.EqualFold(pwd, repwd) {
		this.Data["Message"] = "密码与重复密码不一致"
		this.MobRegister()
		return
	}

	var smsCode string

	key := smsCodePrefix + mobile
	//验证码是否能或取到
	if err := redis_util.Get(key, &smsCode); err != nil {
		this.Data["Message"] = "验证码不正确"
		this.MobRegister()
		return
	}

	if strings.EqualFold(reSmsCode, smsCode) {
		user := entities.User{Cell:mobile, Password:pwd, Nick:mobile, Name:mobile, CreatedAt:time.Now()}
		err := models.NewMobUser(&user)
		if err != nil {
			this.Data["Message"] = "注册失败,请稍后重试"
			this.MobRegister()
			return
		}
		redis_util.Del(key)
		this.SetCurrSession("user", user)
		this.Ctx.Redirect(302, "/")
	}else {
		//验证码不正确
		this.Data["Message"] = "验证码不正确"
		this.MobRegister()
		return
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

