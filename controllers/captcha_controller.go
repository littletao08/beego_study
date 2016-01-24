package controllers

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type CaptchaController struct {
	BaseController
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

func (this *CaptchaController) Get() {
	this.TplName = "login.html"
}

func (this *CaptchaController) Post() {
	this.TplName = "login.html"

	this.Data["Success"] = cpt.VerifyReq(this.Ctx.Request)
}