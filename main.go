package main

import (
	"ascii-art/handler"
	"log"
	"net/http"
)

func main() {
	h, err := handler.NewHandler()
	if err != nil {
		log.Fatalf("failed to initialize handler: %v", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/ascii-art", h.AsciiHandler)
	http.HandleFunc("/about", h.AboutHandler)
	http.HandleFunc("/history", h.HistoryHandler)
	http.HandleFunc("/contact", h.ContactHandler)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
