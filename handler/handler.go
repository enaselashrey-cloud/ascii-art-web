package handler

import (
	"ascii-art/ascii"
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	tmpl *template.Template
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

func NewHandler() (*Handler, error) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Handler{tmpl: tmpl}, nil
}

func (h *Handler) render(
	w http.ResponseWriter,
	status int,
	name string,
	data any,
) {
	var output bytes.Buffer
	if err := h.tmpl.ExecuteTemplate(&output, name, data); err != nil {
		log.Println("template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(output.Bytes()); err != nil {
		log.Println("response write error:", err)
	}
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.render(w, http.StatusMethodNotAllowed, "error", ErrorData{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		return
	}

	if r.URL.Path != "/" {
		h.render(w, http.StatusNotFound, "error", ErrorData{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		return
	}

	h.render(w, http.StatusOK, "index", PageData{})
}

func (h *Handler) AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.render(w, http.StatusMethodNotAllowed, "error", ErrorData{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	if text == "" {
		h.render(w, http.StatusOK, "index", PageData{
			Font:  font,
			Color: color,
		})
		return
	}

	if font == "" {
		font = "standard"
	}

	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		h.render(w, http.StatusBadRequest, "error", ErrorData{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid font",
		})
		return
	}

	result, err := ascii.RenderAscii(text, font)
	if err != nil {
		if errors.Is(err, ascii.ErrUnsupportedCharacter) {
			h.render(w, http.StatusBadRequest, "index", PageData{
				Result: "Invalid input",
				Text:   text,
				Font:   font,
				Color:  color,
			})
			return
		}

		if errors.Is(err, ascii.ErrBannerNotFound) {
			h.render(w, http.StatusNotFound, "error", ErrorData{
				StatusCode: http.StatusNotFound,
				Message:    "Banner Not Found",
			})
			return
		}

		h.render(w, http.StatusInternalServerError, "error", ErrorData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		})
		return
	}

	h.render(w, http.StatusOK, "index", PageData{
		Result: result,
		Text:   text,
		Font:   font,
		Color:  color,
	})
}

func (h *Handler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.render(w, http.StatusMethodNotAllowed, "error", ErrorData{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		return
	}

	h.render(w, http.StatusOK, "about", nil)
}

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.render(w, http.StatusMethodNotAllowed, "error", ErrorData{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		return
	}

	h.render(w, http.StatusOK, "history", nil)
}

func (h *Handler) ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.render(w, http.StatusMethodNotAllowed, "error", ErrorData{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		return
	}

	h.render(w, http.StatusOK, "contact", nil)
}
