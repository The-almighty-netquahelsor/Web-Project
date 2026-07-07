package main

import (
	"fmt"
	"log"
	"net/http"
)



func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Welcome to ASCII-ART-WEB</h1>")
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "", 400)
		return
	}
	text   := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == ""{
		http.Error(w, "", 400)
		return
	}

	
	fmt.Fprintf(w, "Text: %s | Banner: %s", text, banner)

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	fmt.Println("Serving at port: localhost:8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
