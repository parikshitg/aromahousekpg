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

	Debug(c.Data["Gallery"])
}
