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

	page := content.GetPage("en", "reservation")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}
