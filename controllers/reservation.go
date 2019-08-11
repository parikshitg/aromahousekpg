package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type ReservationController struct {
	beego.Controller
}

func (c *ReservationController) Get() {
	c.TplName = "page/reservation.tpl"

	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)

}
