package controllers

import (
	"github.com/astaxie/beego"
	"github.com/parikshitg/aromahousekpg/content"
)

type AccomodationController struct {
	beego.Controller
}

func (c *AccomodationController) Get() {
	c.TplName = "page/accomodation.tpl"

	c.Data["Rooms"], _, _ = content.RoomList("en", "weight", -1, 0)

	page := content.GetPage("en", "accomodation")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta
}
