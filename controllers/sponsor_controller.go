package controllers
import "beego_study/models"

type SponsorController struct {
	BaseController
}

func (c *SponsorController) New() {
	var order = models.Pay(float32(1))
	c.Data["json"] = order
	c.ServeJson()
}