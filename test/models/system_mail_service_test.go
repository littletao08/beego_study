package services

import (
	"testing"
	"fmt"
	"beego_study/services"
)

func TestSendHtmlMail(t *testing.T) {

	content := fmt.Sprintf(services.MAIL_CAPTCHA_TEMPLATE, "11111", "22222")

	services.SendHtmlMail("邮件测试", content, []string{"406504302@qq.com"})
}