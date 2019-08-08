package routers

import (
	"git.urantiatech.com/homestay/aromahousekpg/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/accomodation", &controllers.AccomodationController{})
	beego.Router("/room/:id", &controllers.RoomController{})
	beego.Router("/reservation", &controllers.ReservationController{})
	beego.Router("/restaurant/?:slug", &controllers.RestaurantController{})
	beego.Router("/package", &controllers.PackageController{})
	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/gallery", &controllers.GalleryController{})
	beego.Router("/testimonial", &controllers.TestimonialController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/attraction", &controllers.AttractionController{})
	beego.Router("/activity", &controllers.ActivityController{})
	// beego.ErrorController(&controllers.ErrorController{})

	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/drive", "drive")

}
