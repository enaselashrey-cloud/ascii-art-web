package main

import (
	"ascii-art/handler"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii", handler.AsciiHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/history",handler.History)
	http.HandleFunc("/contact",handler.Contact)
	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
