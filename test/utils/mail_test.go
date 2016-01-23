package utils

import (
	"testing"
	"beego_study/utils"
	"fmt"
)

var smtConfig utils.SmtpConfig
var to = "406504302@qq.com"
func init()  {
	from := "noreply@threeperson.com"
	password := "Threeperson2015"
	host:= "smtp.mxhichina.com"
	addr := "smtp.mxhichina.com:25"

	smtConfig.Addr=addr
	smtConfig.Host=host
	smtConfig.Password=password
	smtConfig.Username=from
}

func TestSendMail(t *testing.T) {

	subject := "hehe"

	content := `
<html>
<body>
<div>欢迎访问Threeperson博客</div>
<div>您的验证码是12345</div>
</body>
</html>`

	err := utils.SendMail(subject, content, smtConfig.Username, []string{to}, smtConfig, true)

	fmt.Println("err:", err)
}


func TestSendTextMail(t *testing.T) {

	subject := "afafd"

	content := "天气很好"

	err := utils.SendMail(subject, content, smtConfig.Username, []string{to}, smtConfig, true)

	fmt.Println("err:", err)
}

