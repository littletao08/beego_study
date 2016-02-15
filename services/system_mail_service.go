package services

import (
	"beego_study/entities"
	"beego_study/db"
	"encoding/json"
	"github.com/astaxie/beego/utils"
)

const MAIL_CAPTCHA_TEMPLATE =
`<div id="mailContentContainer" class="qmbox qm_con_body_content">
	<div>亲爱的Threeperson用户<font color="#f77616"><b></b></font>，</div>
	<br>您好！您的验证码是：<span style="border-bottom-width: 1px; border-bottom-style: dashed; border-bottom-color: rgb(204, 204, 204); z-index: 1; position: static;" t="7" onclick="return false;" data="%v">%v</span>
	<br>本邮件是系统自动发送的，请勿直接回复！感谢您的访问，
	<br>祝您使用愉快！<br>
	<br>Threeperson博客<br>
	<a href="http://www.threeperson.com" target="_blank">www.threeperson.com</a>
	<br>
	<div>
	</div>
	<style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style>
	</div>`

func ValidSystemMail() (entities.SystemMail, error) {

	var mail entities.SystemMail

	db := db.NewDB()

	err := db.QueryTable("system_mail").Filter("valid", entities.SYSTEM_MAIL_VALID_YES).One(&mail)

	return mail, err
}

func mailConfig() (string, error) {
	systemMail, err := ValidSystemMail()

	if nil != err {
		return "", err
	}

	bytes, err := json.Marshal(systemMail)
	if nil != err {
		return "", err
	}

	config := string(bytes)

	return config, err
}

func SendHtmlMail(subject string, content string, to []string) {

	config, err := mailConfig()
	if nil != err {
		return
	}

	mail := utils.NewEMail(config)
	mail.Subject = subject
	mail.HTML = content
	mail.To = to
	mail.Send()
}

func SendTextMail(subject string, content string, to []string) {

	config, err := mailConfig()
	if nil != err {
		return
	}

	mail := utils.NewEMail(config)
	mail.Subject = subject
	mail.Text = content
	mail.To = to

	mail.Send()
}
