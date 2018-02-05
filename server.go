package main

import (
	"myGrpc/controllers"
	"myGrpc/config"
	"net/http"
	"fmt"
	"myGrpc/templmanager"
)

var cfg = config.AppConfig

func init() {
	err := templmanager.LoadTemplates(cfg)

	if err != nil {
		fmt.Println("Could not load templates: ", err)
	}

		http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("web/"))))
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/"))))
		http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("node_modules/"))))

		fmt.Println("Static files served...")

}

func main() {
	server := http.Server{
		Addr: cfg.Host + ":" + cfg.Port,
	}

	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/upload-file", controllers.UploadPage)
	http.HandleFunc("/upload", controllers.UploadHandler)
	http.HandleFunc("/view", controllers.ViewPage)

	server.ListenAndServe()
}
