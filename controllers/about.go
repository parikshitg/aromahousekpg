package controllers

import (
	"github.com/astaxie/beego"
	"github.com/parikshitg/aromahousekpg/content"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.TplName = "page/about.tpl"

	page := content.GetPage("en", "about")
	c.Data["About"] = page

	c.Data["Chef"] = content.GetPage("en", "chef")

	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta
}
