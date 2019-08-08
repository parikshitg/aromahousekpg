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

}
