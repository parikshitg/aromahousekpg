package controllers

import (
	"github.com/astaxie/beego"
	"github.com/parikshitg/aromahousekpg/content"
)

type AttractionController struct {
	beego.Controller
}

func (c *AttractionController) Get() {
	c.TplName = "page/attractions.tpl"

	c.Data["Rooms"], _, _ = content.RoomList("en", "", -1, 0)
	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)

	c.Data["Attractions"], _, _ = content.AttractionList("en", "", -1, 0)

	page := content.GetPage("en", "attraction")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
