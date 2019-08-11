package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type PackageController struct {
	beego.Controller
}

func (c *PackageController) Get() {
	c.TplName = "page/package.tpl"

	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)

	page := content.GetPage("en", "packages")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
