package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.TplName = "page/about.tpl"

	c.Data["About"] = content.GetPage("en", "about-aroma-house")

	c.Data["Chef"] = content.GetPage("en", "chef-bikas-lama")

}
