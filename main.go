package main

import (
	"ascii-art/handler"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
