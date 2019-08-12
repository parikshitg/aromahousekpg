package main

import (
	"log"
	"strings"

	cloudcms "git.urantiatech.com/cloudcms/lightcms"
	s "git.urantiatech.com/cloudcms/lightcms/service"
	_ "git.urantiatech.com/homestay/aromahousekpg/content"
	_ "git.urantiatech.com/homestay/aromahousekpg/routers"
	"git.urantiatech.com/homestay/aromahousekpg/views"
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

	cmsport, err := beego.AppConfig.Int("cmsport")
	if err != nil {
		cmsport = 9090
	}
	go cloudcms.Run(cmsport)

	beego.AddFuncMap("htmlString", views.HTMLString)
	beego.AddFuncMap("removeDiv", views.RemoveDiv)
	beego.AddFuncMap("splitString", strings.Split)
	beego.AddFuncMap("isEven", views.IsEven)
	beego.AddFuncMap("isOdd", views.IsOdd)
	beego.AddFuncMap("firstChar", views.FirstChar)
	beego.AddFuncMap("image", views.Image)
	beego.Run()
}
