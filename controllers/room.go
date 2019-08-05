package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type RoomController struct {
	beego.Controller
}

func (c *RoomController) Get() {
	c.TplName = "page/room.tpl"

}
