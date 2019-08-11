package controllers

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/muesli/sticker"
	"golang.org/x/image/font/gofont/gobold"
)

type PlaceholderController struct {
	beego.Controller
}

func (c *PlaceholderController) Get() {
	size := strings.Split(c.Ctx.Input.Param(":size"), "x")
	if len(size) != 2 {
		c.CustomAbort(http.StatusBadRequest, "Invalid image size")
		return
	}

	width, err1 := strconv.ParseInt(size[0], 10, 32)
	height, err2 := strconv.ParseInt(size[1], 10, 32)
	if err1 != nil || err2 != nil {
		c.CustomAbort(http.StatusBadRequest, "Invalid height and width")
		return
	}

	if width > 4000 || height > 4000 {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	text := c.GetString("text")
	log.Println("text:", text)
	if text == "" {
		text = fmt.Sprintf("%dx%d", width, height)
	}

	gen, err := sticker.NewImageGenerator(sticker.Options{
		TTF:         gobold.TTF,
		MarginRatio: 0.5,
		Foreground:  color.RGBA{0x96, 0x96, 0x96, 0xff},
		Background:  color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	})
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
		return
	}

	img, err := gen.NewPlaceholder(text, int(width), int(height))
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
		return
	}

	c.Ctx.ResponseWriter.Header().Set("Content-Type", "image/png")
	png.Encode(c.Ctx.ResponseWriter, img)
}
