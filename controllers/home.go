package controllers

import (
	//"log"

	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	"golang.org/x/text/language"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "page/home.tpl"

	blogs, total, err := content.BlogList(language.English.String(), "-id", -1, 0)
	if err != nil {
		c.Data["Error"] = err
	}

	/*var b = &content.Blog{
		Title: "Creating blog inside controller",
	}

	if err := b.Create("en", b.Title); err != nil {
		log.Println(err)
	}*/

	c.Data["Blogs"] = blogs
	c.Data["Total"] = total
}
