package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.TplName = "page/about.tpl"

}
