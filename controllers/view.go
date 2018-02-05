package controllers

import (
	"net/http"
	"myGrpc/flashes"
	"myGrpc/templmanager"
	"log"
)

type viewPageData struct {
	Title   string
	Flash   string
	Content string
}

func ViewPage(w http.ResponseWriter, r *http.Request) {
	var err error
	var flash string

	flash, _ = flashes.GetFlash(w, r, "message")

	data := &viewPageData{
		Title:   "Home Page",
		Flash:   flash,
		Content: "This is the page content",
	}

	err = templmanager.RenderTemplate(w, "view.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}