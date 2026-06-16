package handler

import (
	"ascii-art/ascii"
	"errors"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
	Text   string
	Font   string
	Color  string
	Error  string
}

var (
	indexTmpl *template.Template
	aboutTmpl *template.Template
)

func InitTemplates() error {
	var err error

	indexTmpl, err = template.ParseFiles(
		"templates/layout.html",
		"templates/index.html",
	)
	if err != nil {
		return err
	}

	aboutTmpl, err = template.ParseFiles(
		"templates/layout.html",
		"templates/about.html",
	)
	if err != nil {
		return err
	}

	return nil
}

func renderTemplate(
	w http.ResponseWriter,
	tmpl *template.Template,
	data any,
) {
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(
			w,
			"Template Error",
			http.StatusInternalServerError,
		)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, indexTmpl, PageData{})
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
		http.Error(w, "text is required", http.StatusBadRequest)
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
		renderTemplate(w, indexTmpl, PageData{
			Result: "Invalid input",
			Text:   text,
			Font:   font,
			Color:  color,
		})
		return
	}

	renderTemplate(w, indexTmpl, PageData{
		Result: result,
		Text:   text,
		Font:   font,
		Color:  color,
	})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, aboutTmpl, nil)
}
