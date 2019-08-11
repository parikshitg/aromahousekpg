package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	c.TplName = "page/blog.tpl"

	page := content.GetPage("en", "blog")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
