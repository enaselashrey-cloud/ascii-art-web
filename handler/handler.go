package handler

import (
	"ascii-art/ascii"
	"errors"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	indexTmpl   *template.Template
	aboutTmpl   *template.Template
	historyTmpl *template.Template
	contactTmpl *template.Template
	errorTmpl   *template.Template
}

func NewHandler() (*Handler, error) {
	indexTmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/index.html",
	)
	if err != nil {
		return nil, err
	}

	aboutTmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/about.html",
	)
	if err != nil {
		return nil, err
	}

	historyTmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/history.html",
	)
	if err != nil {
		return nil, err
	}

	contactTmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/contact.html",
	)
	if err != nil {
		return nil, err
	}

	errorTmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/error.html",
	)
	if err != nil {
		return nil, err
	}

	return &Handler{
		indexTmpl:   indexTmpl,
		aboutTmpl:   aboutTmpl,
		historyTmpl: historyTmpl,
		contactTmpl: contactTmpl,
		errorTmpl:   errorTmpl,
	}, nil
}

type PageData struct {
	Result string
	Text   string
	Font   string
	Color  string
}

type ErrorData struct {
	StatusCode int
	Message    string
}

func renderTemplate(
	w http.ResponseWriter,
	tmpl *template.Template,
	data any,
) {
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println("template execution error:", err)
	}
}

func (h *Handler) renderError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	data := ErrorData{
		StatusCode: status,
		Message:    message,
	}

	w.WriteHeader(status)
	renderTemplate(w, h.errorTmpl, data)
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	if r.URL.Path != "/" {
		h.renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	renderTemplate(w, h.indexTmpl, PageData{})
}

func (h *Handler) AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
		renderTemplate(w, h.indexTmpl, PageData{})
		return
	}

	if font == "" {
		font = "standard"
	}

	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		h.renderError(w, http.StatusBadRequest, "Invalid font")
		return
	}

	banner, err := ascii.LoadBanner(font)
	if err != nil {
		if errors.Is(err, ascii.ErrBannerNotFound) {
			h.renderError(w, http.StatusNotFound, "Font not found")
			return
		}

		h.renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	result, err := ascii.RenderAscii(text, banner)
	if err != nil {
		if errors.Is(err, ascii.ErrInvalidBanner) {
			h.renderError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		renderTemplate(w, h.indexTmpl, PageData{
			Result: "Invalid input",
			Text:   text,
			Font:   font,
			Color:  color,
		})
		return
	}

	renderTemplate(w, h.indexTmpl, PageData{
		Result: result,
		Text:   text,
		Font:   font,
		Color:  color,
	})
}

func (h *Handler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	renderTemplate(w, h.aboutTmpl, nil)
}

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	renderTemplate(w, h.historyTmpl, nil)
}

func (h *Handler) ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	renderTemplate(w, h.contactTmpl, nil)
}
