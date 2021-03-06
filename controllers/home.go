package controllers

import (
	"log"
	"myGrpc/flashes"
	"myGrpc/templmanager"
	"net/http"
)

type homePageData struct {
	Title   string
	Flash   string
	Content string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var err error
	var flash string

	flash, _ = flashes.GetFlash(w, r, "message")

	data := &homePageData{
		Title:   "Home Page",
		Flash:   flash,
		Content: "This is the page content",
	}

	err = templmanager.RenderTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}
