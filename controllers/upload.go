package controllers

import (
	"net/http"
	"myGrpc/flashes"
	"myGrpc/templmanager"
	"log"
)

type uploadPageData struct {
	Title   string
	Flash   string
	Content string
}

func UploadPage(w http.ResponseWriter, r *http.Request) {
	var err error
	var flash string

	flash, _ = flashes.GetFlash(w, r, "message")

	data := &uploadPageData{
		Title:   "Home Page",
		Flash:   flash,
		Content: "This is the page content",
	}

	err = templmanager.RenderTemplate(w, "upload.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}