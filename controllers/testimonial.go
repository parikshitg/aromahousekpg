package controllers

import (
	"github.com/astaxie/beego"
	"github.com/parikshitg/aromahousekpg/content"
)

type TestimonialController struct {
	beego.Controller
}

func (c *TestimonialController) Get() {
	c.TplName = "page/testimonial.tpl"

	c.Data["Testimonials"], _, _ = content.TestimonialList("en", "", -1, 0)

	page := content.GetPage("en", "testimonials")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
