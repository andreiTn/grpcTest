package main

import (
	"net/http"
	"myGrpc/controllers"
)

func main()  {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/upload", controllers.FileHandler)
	http.HandleFunc("/show-image", controllers.ShowImage)

	http.ListenAndServe(":8080", nil)
}