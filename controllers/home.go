package controllers

import (
	"net/http"
	"myGrpc/flashes"
	"myGrpc/services"
	"fmt"
)
const InputName = "uploadImage"

type page struct {
	Title string
	InputName string
	Flash string
}

func Home(w http.ResponseWriter, r *http.Request)  {
	flash, err := flashes.GetFlash(w, r, "message")

	if err != nil {
		fmt.Println("Flash error: ", err)
	}
	stringFlash := fmt.Sprintf("%s", flash)

	homepage := page{
		Title: "page title",
		InputName: InputName,
		Flash: stringFlash,
	}
	services.Tpl.ExecuteTemplate(w, "index.gohtml", homepage)
}