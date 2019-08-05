package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type ActivityController struct {
	beego.Controller
}

func (c *ActivityController) Get() {
	c.TplName = "page/activity.tpl"

}
