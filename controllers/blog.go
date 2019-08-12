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

	if slug := c.Ctx.Input.Param(":slug"); slug != "" {
		c.Data["Post"] = content.GetBlog("en", slug)
	} else {
		c.Data["BlogList"], _, _ = content.BlogList("en", "-id", -1, 0)
	}

}
