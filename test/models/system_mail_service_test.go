package models

import (
	"testing"
	"fmt"
	"beego_study/models"
)

func TestSendHtmlMail(t *testing.T) {

	content := fmt.Sprintf(models.MAIL_CAPTCHA_TEMPLATE, "11111", "22222")

	models.SendHtmlMail("邮件测试", content, []string{"406504302@qq.com"})
}