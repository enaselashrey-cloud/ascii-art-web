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

tmpl, err := template.ParseFiles(
    "templates/index.html",
    "templates/navbar.html",
)

if err != nil {
    http.Error(w, err.Error(), 500)
    return
}
	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
    err = tmpl.Execute(w, PageData{})
    if err != nil {
        http.Error(w, err.Error(), 500)
    }
    return
}

	if font == "" {
		font = "standard"
	}

	banner, err := ascii.NewBanner(font)
if err != nil {
    http.Error(w, err.Error(), 500)
    return
}

	opts := ascii.Options{
		Text: text,
	}

	result, err := ascii.Render(banner, opts)
if err != nil {
    http.Error(w, err.Error(), 500)
    return
}
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
