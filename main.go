package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var index *template.Template
var new *template.Template

func init() {
	index = template.Must(template.ParseFiles("index.html", "layout.html"))
	new = template.Must(template.ParseFiles("new.html", "layout.html"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/contacts/new", newHandler)
	r.HandleFunc("/contacts", createHandler).Methods("POST")

	port := ":8080"
	fmt.Printf("Serving from port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	index.ExecuteTemplate(w, "layout", "")
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	new.ExecuteTemplate(w, "layout", "")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(body))
}
