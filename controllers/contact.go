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

type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.TplName = "page/contact.tpl"

	page := content.GetPage("en", "contact")
	c.Data["Contact"] = page

	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

	c.Data["ErrrorMsg"] = "Cannot contact"
}

func (c *ContactController) Post() {
	c.TplName = "page/contact.tpl"

	page := content.GetPage("en", "contact")
	c.Data["Contact"] = page

	name := c.GetString("name")
	email := c.GetString("email")
	phone := c.GetString("phone")
	message := c.GetString("message")

	subject := fmt.Sprintf("ContactForm: Message from %s (phone: %s)", name, phone)
	body := message

	// Send mail
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
		c.Data["Message"] = message
	} else {
		c.Data["SuccessMsg"] = "Your message has been sent"
	}

}
