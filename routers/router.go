package routers

import (
	"git.urantiatech.com/homestay/aromahousekpg/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
