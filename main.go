package main

import (
	"ascii-color/ascii"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r*http.Request){
text := r.FormValue("text")// هات قيمة انبوت من صغحة html 
font := r.FormValue("font")

if text == "" {
		http.ServeFile(w, r, "templates/index.html")
		return
	}

if font == "" {
	font = "standard"
}
	

banner, _:= ascii.NewBanner(font)
opts := ascii.Options{
	Text : text,
}
result, _:= ascii.Render(banner,opts)
fmt.Fprintln(w,result)

}


func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}