package controllers

import (
	//"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
	//"golang.org/x/text/language"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	c.TplName = "page/blog.tpl"

}
