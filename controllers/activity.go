package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type ActivityController struct {
	beego.Controller
}

func (c *ActivityController) Get() {
	c.TplName = "page/activities.tpl"

	c.Data["Rooms"], _, _ = content.RoomList("en", "", -1, 0)
	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)

	c.Data["Activities"], _, _ = content.ActivityList("en", "", -1, 0)

}
