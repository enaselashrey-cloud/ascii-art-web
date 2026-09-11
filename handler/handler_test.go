package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func chdir(t *testing.T, dir string) {
	t.Helper()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() returned an error: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("os.Chdir(%q) returned an error: %v", dir, err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})
}

func newTestHandler(t *testing.T) *Handler {
	t.Helper()
	chdir(t, "..")

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("NewHandler() returned an error: %v", err)
	}
	return h
}

func TestPageHandlers(t *testing.T) {
	h := newTestHandler(t)

	tests := []struct {
		name       string
		handler    http.HandlerFunc
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{name: "home", handler: h.HomeHandler, method: http.MethodGet, path: "/", wantStatus: http.StatusOK, wantBody: "Enter your text"},
		{name: "not found", handler: h.HomeHandler, method: http.MethodGet, path: "/missing", wantStatus: http.StatusNotFound, wantBody: "Page Not Found"},
		{name: "home wrong method", handler: h.HomeHandler, method: http.MethodPost, path: "/", wantStatus: http.StatusMethodNotAllowed, wantBody: "Method Not Allowed"},
		{name: "about", handler: h.AboutHandler, method: http.MethodGet, path: "/about", wantStatus: http.StatusOK, wantBody: "About The Project"},
		{name: "history", handler: h.HistoryHandler, method: http.MethodGet, path: "/history", wantStatus: http.StatusOK, wantBody: "History of ASCII Art"},
		{name: "contact", handler: h.ContactHandler, method: http.MethodGet, path: "/contact", wantStatus: http.StatusOK, wantBody: "Contact Us"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			recorder := httptest.NewRecorder()

			tt.handler(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Errorf("status = %d; want %d", recorder.Code, tt.wantStatus)
			}
			if !strings.Contains(recorder.Body.String(), tt.wantBody) {
				t.Errorf("body does not contain %q", tt.wantBody)
			}
		})
	}
}

func TestAsciiHandler(t *testing.T) {
	h := newTestHandler(t)

	tests := []struct {
		name       string
		method     string
		form       url.Values
		wantStatus int
		wantBody   string
	}{
		{name: "successful generation", method: http.MethodPost, form: url.Values{"text": {"A"}, "font": {"standard"}}, wantStatus: http.StatusOK, wantBody: "/_/    \\_\\"},
		{name: "default font", method: http.MethodPost, form: url.Values{"text": {"A"}}, wantStatus: http.StatusOK, wantBody: "/_/    \\_\\"},
		{name: "empty text", method: http.MethodPost, form: url.Values{}, wantStatus: http.StatusOK, wantBody: "Enter your text"},
		{name: "invalid font", method: http.MethodPost, form: url.Values{"text": {"A"}, "font": {"unknown"}}, wantStatus: http.StatusBadRequest, wantBody: "Invalid font"},
		{name: "unsupported character", method: http.MethodPost, form: url.Values{"text": {"é"}, "font": {"standard"}}, wantStatus: http.StatusBadRequest, wantBody: "Invalid input"},
		{name: "wrong method", method: http.MethodGet, form: nil, wantStatus: http.StatusMethodNotAllowed, wantBody: "Method Not Allowed"},
		{name: "HTML is escaped", method: http.MethodPost, form: url.Values{"text": {"<script>"}, "font": {"standard"}}, wantStatus: http.StatusOK, wantBody: "&lt;script&gt;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body *strings.Reader
			if tt.form == nil {
				body = strings.NewReader("")
			} else {
				body = strings.NewReader(tt.form.Encode())
			}
			req := httptest.NewRequest(tt.method, "/ascii-art", body)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			recorder := httptest.NewRecorder()

			h.AsciiHandler(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Errorf("status = %d; want %d", recorder.Code, tt.wantStatus)
			}
			if !strings.Contains(recorder.Body.String(), tt.wantBody) {
				t.Errorf("body does not contain %q", tt.wantBody)
			}
		})
	}
}

func TestAsciiHandlerMissingSupportedBanner(t *testing.T) {
	h := newTestHandler(t)
	chdir(t, t.TempDir())

	form := url.Values{"text": {"A"}, "font": {"standard"}}
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	h.AsciiHandler(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("status = %d; want %d", recorder.Code, http.StatusNotFound)
	}
}

func TestRenderExecutionError(t *testing.T) {
	tmpl := template.Must(template.New("broken").Parse(
		`{{define "broken"}}partial{{template "missing" .}}{{end}}`,
	))
	h := &Handler{tmpl: tmpl}
	recorder := httptest.NewRecorder()

	h.render(recorder, http.StatusOK, "broken", nil)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("status = %d; want %d", recorder.Code, http.StatusInternalServerError)
	}
	if got, want := recorder.Body.String(), "Internal Server Error\n"; got != want {
		t.Errorf("body = %q; want %q", got, want)
	}
}
