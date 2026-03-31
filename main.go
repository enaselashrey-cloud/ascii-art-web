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
	"templates/layout.html",
	 "templates/navbar.html",
    "templates/index.html",
)

if err != nil {
    http.Error(w, err.Error(), 500)
    return
}
	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
    err = tmpl.ExecuteTemplate(w, "layout", PageData{})
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

	tmpl.ExecuteTemplate(w, "layout", PageData{Result: template.HTML(result),
	Text:   text,
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/navbar.html",
		"templates/about.html", // 👈 الصفحة الجديدة
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/style.css", http.FileServer(http.Dir("templates")))

	http.HandleFunc("/", handler)
	http.HandleFunc("/about", aboutHandler) // 👈 دي أهم سطر

	http.ListenAndServe(":8080", nil)
}
