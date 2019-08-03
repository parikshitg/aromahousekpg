package main

import (
	cloudcms "git.urantiatech.com/cloudcms/lightcms"
	_ "git.urantiatech.com/homestay/aromahousekpg/content"
	_ "git.urantiatech.com/homestay/aromahousekpg/routers"
	"github.com/astaxie/beego"
	"golang.org/x/text/language"
)

func main() {

	var languages = []language.Tag{
		language.English,
		language.French,
		language.Spanish,
		language.Russian,
	}

	cloudcms.Languages(languages)

	cmsport, err := beego.AppConfig.Int("cmsport")
	if err != nil {
		cmsport = 9090
	}
	go cloudcms.Run(cmsport)

	beego.Run()

}
