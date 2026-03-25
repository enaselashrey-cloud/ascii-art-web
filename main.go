package main

import (
	"ascii-color/ascii"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r*http.Request){
text := r.FormValue("text")// هات قيمة انبوت من صغحة html 
if text == "" {
		http.ServeFile(w, r, "templates/index.html")
		return
	}

banner, _:= ascii.NewBanner("standard")
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