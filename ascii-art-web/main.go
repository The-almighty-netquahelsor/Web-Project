package main

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/ascii-art"
)

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHanler)

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("fatal: server failed:", err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 mehtod not allowed", 405)
		return
	}
	renderPage(w, PageData{Banner: "standard"})
}

func asciiArtHanler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", 405)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 bad request", 400)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		renderPage(w, PageData{
			Input:  text,
			Banner: banner,
			Error:  "please enter some text and choose a banner.",
		})
		return
	}
	bannerLines, err := asciiart.LoadBanner(banner)
	if err != nil {
		http.Error(w, "404 not: "+err.Error(), 404)
		return
	}
	result := asciiart.Render(text, bannerLines)

	renderPage(w, PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	})
}

func renderPage(w http.ResponseWriter, data PageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 internal server error", 500)
		return
	}
}
