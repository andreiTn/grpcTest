package services

import (
	"net/http"
	"fmt"
)

func init() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("web/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("node_modules/"))))

	fmt.Println("Static files served...")
}