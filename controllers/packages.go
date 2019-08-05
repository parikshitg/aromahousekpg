package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type PackageController struct {
	beego.Controller
}

func (c *PackageController) Get() {
	c.TplName = "page/package.tpl"

}
