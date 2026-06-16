package main

import (
	"ascii-art/handler"
	"log"
	"net/http"
)

func main() {
	// Initialize templates at startup
	if err := handler.InitTemplates(); err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii", handler.AsciiArtHandler)
	http.HandleFunc("/about", handler.AboutHandler)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
