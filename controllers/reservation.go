package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type ReservationController struct {
	beego.Controller
}

func (c *ReservationController) Get() {
	c.TplName = "page/reservation.tpl"

}
