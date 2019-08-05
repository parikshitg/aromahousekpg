package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type RestaurantController struct {
	beego.Controller
}

func (c *RestaurantController) Get() {
	c.TplName = "page/restaurant.tpl"

}
