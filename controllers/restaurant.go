package controllers

import (
	"strings"

	"git.urantiatech.com/homestay/aromahousekpg/content"
	"github.com/astaxie/beego"
)

type RestaurantController struct {
	beego.Controller
}

func (c *RestaurantController) Get() {
	c.TplName = "page/restaurant.tpl"
	slug := c.Ctx.Input.Param(":slug")

	page := content.GetPage("en", "restaurant")
	c.Data["Restaurant"] = page

	meta := make(map[string]string)
	meta["keywords"] = page.MetaKeywords
	meta["description"] = page.MetaDescription
	c.Data["Meta"] = meta

	var categories = make(map[string]string)

	// Food Menu
	var foodmenu = make(map[string](map[string]content.Food))

	foodlist, _, _ := content.FoodList("en", "", -1, 0)
	for _, v := range foodlist {
		if _, ok := foodmenu[v.Cuisine]; !ok {
			foodmenu[v.Cuisine] = make(map[string]content.Food)
		}
		cuisine := foodmenu[v.Cuisine]
		cuisine[v.Title] = v
	}

	for _, v := range foodmenu {
		for _, v1 := range v {
			categories[strings.ToLower(v1.Category)] = v1.Category
		}
	}

	// Drink Menu
	var drinkmenu = make(map[string](map[string]content.Drink))

	drinklist, _, _ := content.DrinkList("en", "", -1, 0)
	for _, v := range drinklist {
		if _, ok := drinkmenu[v.Category]; !ok {
			drinkmenu[v.Category] = make(map[string]content.Drink)
		}
		category := drinkmenu[v.Category]
		category[v.Title] = v
	}

	for _, v := range drinkmenu {
		for _, v1 := range v {
			categories[strings.ToLower(v1.Category)] = v1.Category
		}
	}

	c.Data["Categories"] = categories

	if slug == "" {
		c.Data["FoodMenu"] = foodmenu
		c.Data["DrinkMenu"] = drinkmenu
		return
	}

	// Filter foods by category
	for k, v := range foodmenu {
		submenu := v
		for k1, v1 := range v {
			if strings.ToLower(v1.Category) != slug {
				delete(submenu, k1)
			}
		}
		if len(submenu) == 0 {
			delete(foodmenu, k)
		}
	}

	c.Data["FoodMenu"] = foodmenu

	// Filter drinks by category
	for k, v := range drinkmenu {
		submenu := v
		for k1, v1 := range v {
			if strings.ToLower(v1.Category) != slug {
				delete(submenu, k1)
			}
		}
		if len(submenu) == 0 {
			delete(drinkmenu, k)
		}
	}

	c.Data["DrinkMenu"] = drinkmenu

}
