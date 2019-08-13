package controllers

import (
	"math/rand"

	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "page/home.tpl"

	c.Data["Banners"], _, _ = content.BannerList("en", "weight", -1, 0)

	c.Data["Sidebar"] = content.GetPage("en", "why-aroma-house")

	c.Data["Rooms"], _, _ = content.RoomList("en", "weight", -1, 0)

	testimonials, count, _ := content.TestimonialList("en", "", -1, 0)
	if count > 0 {
		c.Data["Testimonial"] = testimonials[rand.Intn(count)]
	}

	page := content.GetPage("en", "home")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
