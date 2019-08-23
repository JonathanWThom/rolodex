package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var index *template.Template

func init() {
	index = template.Must(template.ParseFiles("index.html"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)

	port := ":8080"
	fmt.Printf("Serving from port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	index.Execute(w, nil)
}
