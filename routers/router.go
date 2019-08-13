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
	beego.Router("/packages", &controllers.PackageController{})
	beego.Router("/blog/?:slug", &controllers.BlogController{})
	beego.Router("/gallery", &controllers.GalleryController{})
	beego.Router("/testimonials", &controllers.TestimonialController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/attractions", &controllers.AttractionController{})
	beego.Router("/activities", &controllers.ActivityController{})
	beego.Router("/placeholder/:size([0-9]+x[0-9]+).png", &controllers.PlaceholderController{})

	// beego.ErrorController(&controllers.ErrorController{})

	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/drive", "drive")

}
