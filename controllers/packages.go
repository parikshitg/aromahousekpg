package controllers

import (
	"github.com/astaxie/beego"
	"github.com/parikshitg/aromahousekpg/content"
)

type PackageController struct {
	beego.Controller
}

func (c *PackageController) Get() {
	c.TplName = "page/package.tpl"

	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)
	c.Data["Activities"], _, _ = content.ActivityList("en", "", -1, 0)
	c.Data["Attractions"], _, _ = content.AttractionList("en", "", -1, 0)

	page := content.GetPage("en", "packages")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
