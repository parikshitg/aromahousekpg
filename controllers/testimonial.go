package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type TestimonialController struct {
	beego.Controller
}

func (c *TestimonialController) Get() {
	c.TplName = "page/testimonial.tpl"

}
