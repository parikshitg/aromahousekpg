package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type GalleryController struct {
	beego.Controller
}

func (c *GalleryController) Get() {
	c.TplName = "page/gallery.tpl"

}
