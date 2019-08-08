package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type AccomodationController struct {
	beego.Controller
}

func (c *AccomodationController) Get() {
	c.TplName = "page/accomodation.tpl"

	c.Data["Rooms"], _, _ = content.RoomList("en", "weight", -1, 0)
}
