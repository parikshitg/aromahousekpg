package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.TplName = "page/contact.tpl"

	page := content.GetPage("en", "contact")
	c.Data["Contact"] = page

	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
