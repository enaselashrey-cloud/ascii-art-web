package handler

import (
	"ascii-art/ascii"
	"errors"
	"html/template"
	"net/http"
)

type PageData struct {
	Result template.HTML
	Text   string
	Font   string 
	Color  string 
}

func renderTemplate(
	w http.ResponseWriter,
	page string,
	data any,
) {

	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/"+page,
	))

	err := tmpl.ExecuteTemplate(w, "layout", data)

	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
		renderTemplate(w, "index.html", PageData{})
		return
	}

	if font == "" {
		font = "standard"
	}

	banner, err := ascii.LoadBanner(font)
	if err != nil {
		if errors.Is(err, ascii.ErrBannerNotFound) {
			http.Error(w, "Font not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	result, err := ascii.RenderAscii(text, banner)
	if err != nil {
		if errors.Is(err, ascii.ErrInvalidBanner) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		renderTemplate(w, "index.html", PageData{
			Result: template.HTML("Invalid input"),
			Text:   text,
		})
		return
	}

	if color != "" {
		result = "<span style='color:" + color + "'>" + result + "</span>"
	}

	renderTemplate(w, "index.html", PageData{
		Result: template.HTML(result),
		Text:   text,
		Font:   font, 
		Color:  color,
	})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func History(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "history.html",nil)
}

func Contact(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "contact.html",nil)
}
