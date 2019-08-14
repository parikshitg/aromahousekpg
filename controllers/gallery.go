package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type GalleryController struct {
	beego.Controller
}

func (c *GalleryController) Get() {
	c.TplName = "page/gallery.tpl"

	c.Data["Gallery"], _, _ = content.GalleryList("en", "", -1, 0)

	page := content.GetPage("en", "gallery")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
