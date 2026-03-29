package main

import (
	"ascii-color/ascii"
	"html/template"

	"net/http"
)

type PageData struct {
	Result template.HTML
	Text   string
}

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
		tmpl.Execute(w, nil)
		return
	}

	if font == "" {
		font = "standard"
	}

	banner, _ := ascii.NewBanner(font)

	opts := ascii.Options{
		Text: text,
	}

	result, _ := ascii.Render(banner, opts)
	if color != "" {
		result = "<span style='color:" + color + "'>" + result + "</span>"
	}

	tmpl.Execute(w, PageData{Result: template.HTML(result),
	Text:   text,
	})
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/style.css", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
