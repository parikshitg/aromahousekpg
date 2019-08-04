package controllers

import (
	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	"golang.org/x/text/language"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "home.tpl"

	blogs, total, err := content.BlogList(language.English.String(), "-id", -1, 0)
	if err != nil {
		c.Data["Error"] = err
	}

	c.Data["Blogs"] = blogs
	c.Data["Total"] = total
}
