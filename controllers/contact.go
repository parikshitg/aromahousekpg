package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.TplName = "page/contact.tpl"

	c.Data["Contact"] = content.GetPage("en", "contact")

}
