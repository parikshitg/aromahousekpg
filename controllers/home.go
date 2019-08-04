package controllers

import (
	"log"

	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	"golang.org/x/text/language"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "home.tpl"

	blogs, total, err := content.BlogList1(language.English.String(), "-id", -1, 0)
	if err != nil {
		c.Data["Error"] = err
	}
	log.Println("blogs:", blogs, ", total:", total, ", err:", err)
	c.Data["Blogs"] = blogs
	c.Data["Total"] = total
}
