package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type RoomController struct {
	beego.Controller
}

func (c *RoomController) Get() {
	c.TplName = "page/room.tpl"

	if slug := c.Ctx.Input.Param(":id"); slug != "" {
		c.Data["Room"] = content.GetRoom("en", slug)
	}

	page := content.GetPage("en", "home")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
