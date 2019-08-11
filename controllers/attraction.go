package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type AttractionController struct {
	beego.Controller
}

func (c *AttractionController) Get() {
	c.TplName = "page/attractions.tpl"

	c.Data["Rooms"], _, _ = content.RoomList("en", "", -1, 0)
	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)

	c.Data["Attractions"], _, _ = content.AttractionList("en", "", -1, 0)

}
