package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

//Exercise 1
func pingPong(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "pong")
}
//Exercise 2
func QueryParaAndPathValidation(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == ""{
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

//Exercise 3
func TextCounter(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		fmt.Fprint(w, "Send a POST request with text to count words")
	}
	if r.Method == http.MethodPost{
		body, err := io.ReadAll(r.Body)
		if err != nil{
			http.Error(w, "", 404)
		}
		fmt.Fprintf(w, "%d", len(body))
	}
}

//Exercise 4
func BasicMathApiMultiQueryPara(w http.ResponseWriter, r *http.Request){
	op := r.URL.Query().Get("op")
	NumA := r.URL.Query().Get("a")
	NumB := r.URL.Query().Get("b")

	valA, err := strconv.Atoi(NumA)
	if err != nil{
		http.Error(w, "", 400)
		return
	}
	valB, err := strconv.Atoi(NumB)
	if err != nil{
		http.Error(w, "", 400)
		return
	}
	var Result int
	switch op {
	case "add":
		Result = valA + valB

	case "subtract":
		Result = valA - valB

	case "multiply":
		Result = valA * valB
	
	default:
		http.Error(w, "", 404)
	}
	fmt.Fprintf(w, "Result: %d", Result)
}

//Exercise 5
func UserAgent(w http.ResponseWriter, r *http.Request){
	header := r.Header.Get("User-Agent")
	if header == ""{
		header = "User-Agent"
	}
	fmt.Fprintf(w, "You are visiting us using: %s", header)
}

//Exercise 6

func SecureDashboard(w http.ResponseWriter, r *http.Request) {
	const key = "secret123"
	Api := r.Header.Get("X-API-Key")

	if Api != key{
		http.Error(w, "", 401)
		return
	}
	fmt.Fprint(w, "Welcome")
}

//Exercise 7
func SimpleRedirector(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}
func newVersionHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to version 2")
}
func main(){
	
	http.HandleFunc("/ping", pingPong)
	http.HandleFunc("/hello", QueryParaAndPathValidation)
	http.HandleFunc("/count", TextCounter)
	http.HandleFunc("/calculate", BasicMathApiMultiQueryPara)
	http.HandleFunc("/agent", UserAgent)
	http.HandleFunc("/dashboard", SecureDashboard)
	http.HandleFunc("/legacy", SimpleRedirector)
	http.HandleFunc("/v2", newVersionHandler)
	log.Println("Serving to LocalHost:8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

