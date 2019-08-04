package routers

import (
	"git.urantiatech.com/homestay/aromahousekpg/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	// beego.ErrorController(&controllers.ErrorController{})

	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/drive", "drive")

}
