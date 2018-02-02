package controllers

import (
	"net/http"
	"myGrpc/flashes"
	"fmt"
	"myGrpc/services"
)

type ShowData struct {
	Info interface{}
	Flash string
}

func ShowImage(w http.ResponseWriter, r *http.Request)  {
	var showData ShowData

	if services.ImagePage == nil {
		flash := []byte("No images to display.")
		flashes.SetFlash(w, "message", flash)
		http.Redirect(w, r, "/", 301)
	}

	flash, err := flashes.GetFlash(w, r, "message")
	stringFlash := fmt.Sprintf("%s", flash)

	if err != nil {
		fmt.Println("Flash error: ", err)
	}

	c := fmt.Sprintf("%s", flash)
	fmt.Println(c)

	showData.Info = services.ImagePage
	showData.Flash = stringFlash

	services.Tpl.ExecuteTemplate(w, "image.gohtml", showData)
}
