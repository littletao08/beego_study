package controllers
import "beego_study/services"

type SponsorController struct {
	BaseController
}

func (c *SponsorController) New() {
	var order = services.Pay(float32(1))
	c.Data["json"] = order
	c.ServeJSON()
}