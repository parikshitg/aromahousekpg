package main

import (
	"log"

	s "git.urantiatech.com/cloudcms/cloudcms/service"
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

	dbfile := beego.AppConfig.String("dbfile")
	if dbfile == "" {
		dbfile = "cloudcms.db"
	}

	if err := s.Initialize("db/" + dbfile); err != nil {
		log.Fatal(err.Error())
	}
	if err := s.RebuildIndex(); err != nil {
		log.Fatal(err.Error())
	}

	beego.Run()
}
