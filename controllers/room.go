package controllers

import (
	"github.com/astaxie/beego"
	"github.com/parikshitg/aromahousekpg/content"
)

type RoomController struct {
	beego.Controller
}

func (c *RoomController) Get() {
	c.TplName = "page/room.tpl"

	if slug := c.Ctx.Input.Param(":id"); slug != "" {
		c.Data["Room"] = content.GetRoom("en", slug)
	}

	c.Data["Rooms"], _, _ = content.RoomList("en", "weight", -1, 0)

	page := content.GetPage("en", "home")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
