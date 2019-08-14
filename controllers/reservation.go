package controllers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"git.urantiatech.com/homestay/aromahousekpg/content"
	mailapi "git.urantiatech.com/mail/mail/api"
	"github.com/astaxie/beego"
)

type ReservationController struct {
	beego.Controller
}

func (c *ReservationController) Get() {
	c.TplName = "page/reservation.tpl"

	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)
	c.Data["Rooms"], _, _ = content.RoomList("en", "", -1, 0)

	c.Data["Adults"] = 0
	c.Data["Childs"] = 0

	page := content.GetPage("en", "reservation")
	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

}

func (c *ReservationController) Post() {
	form := c.GetString("form")
	c.Data["Form"] = form
	c.Data["Packages"], _, _ = content.PackageList("en", "", -1, 0)
	c.Data["Rooms"], _, _ = content.RoomList("en", "", -1, 0)

	if form != "check" && form != "reserve" {
		c.Get()
	}

	c.TplName = "page/reservation.tpl"

	if form == "check" {
		c.Data["ArrivalDate"] = c.GetString("arrival-date")
		c.Data["DepartureDate"] = c.GetString("departure-date")
		c.Data["Adults"], _ = c.GetInt("adults")
		c.Data["Childs"], _ = c.GetInt("childs")
		c.Data["RoomType"] = c.GetString("roomtype")
		c.Data["SucessMsg"] = "Rooms are available on the requested date, you many now proceed with the booking."
		return
	}

	// Continue the booking request
	name := c.GetString("name")
	email := c.GetString("email")
	phone := c.GetString("phone")
	arrival := c.GetString("arrival-date")
	departure := c.GetString("departure-date")
	adults, _ := c.GetInt("adults")
	childs, _ := c.GetInt("childs")
	roomtype := c.GetString("roomtype")
	message := c.GetString("message")

	c.Data["Adults"] = 0
	c.Data["Childs"] = 0

	// Send mail

	subject := "ContactForm: " + c.GetString("subject")
	body := fmt.Sprintf("You have received a reservation request from <strong>%s</strong>.\n", name)
	body += fmt.Sprintf("\n")
	body += fmt.Sprintf("Room Type: %v\n", roomtype)
	body += fmt.Sprintf("Arrival Date: %v.\n", arrival)
	body += fmt.Sprintf("Departure Date: %v.\n", departure)
	body += fmt.Sprintf("Number of Adults: %v\n", adults)
	body += fmt.Sprintf("Number of Childs: %v\n", childs)
	body += fmt.Sprintf("Phone number: %v\n", phone)
	body += fmt.Sprintf("\n")
	body += fmt.Sprintf("%s\n", message)

	mail := mailapi.Mail{
		From:    "Aroma House <contact@aromahousekpg.com>",
		To:      "Aroma House <contact@aromahousekpg.com>",
		ReplyTo: fmt.Sprintf("%s <%s>", name, email),
		Subject: subject,
		HTML:    strings.Replace(body, "\n", "<br/>", -1),
	}

	err := mailapi.SendMail(&mail, os.Getenv("MAIL_SVC"))
	if err != nil {
		log.Println(err.Error())
		c.Data["ErrorMsg"] = "Couldn't send your messgae due to some technical problem. Please try again later or call Aroma House."
		c.Data["Name"] = name
		c.Data["Email"] = email
		c.Data["Phone"] = phone
		c.Data["ArrivalDate"] = arrival
		c.Data["DepartureDate"] = departure
		c.Data["Adults"] = adults
		c.Data["Childs"] = childs
		c.Data["RoomType"] = roomtype
		c.Data["Message"] = message
	} else {
		c.Data["SuccessMsg"] = "Your reservation request has been sent"
	}

}
