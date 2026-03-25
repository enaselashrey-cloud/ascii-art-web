package main

import (
	"ascii-color/ascii"
		"html/template"

	"net/http"
)
type PageData struct {
	Result string
}
func handler(w http.ResponseWriter, r *http.Request){

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
		Text:  text,
		Color: color,
	}

	result, _ := ascii.Render(banner, opts)

	tmpl.Execute(w, PageData{Result: result})
}


func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}