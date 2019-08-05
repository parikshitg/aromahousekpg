package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type AttractionController struct {
	beego.Controller
}

func (c *AttractionController) Get() {
	c.TplName = "page/attraction.tpl"

}
