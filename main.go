package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"log"
	"net/http"
)

var index *template.Template
var new *template.Template
var db *gorm.DB
var err error

type Contact struct {
	gorm.Model
	FirstName string
	LastName  string
}

func init() {
	index = template.Must(template.ParseFiles("message.html", "index.html", "layout.html"))
	new = template.Must(template.ParseFiles("new.html", "layout.html"))
}

func main() {
	// "postgresql://postgres:postgres@db/rolodex_development?sslmode=disable"
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=rolodex_development")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Contact{})

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
	r.ParseForm()
	contact := &Contact{FirstName: r.FormValue("firstname"), LastName: r.FormValue("lastname")}
	errors := db.Create(contact).GetErrors()
	if len(errors) != 0 {
		var msg string
		for i := 0; i < len(errors); i++ {
			msg += fmt.Sprintf("%s\n", errors[i].Error())
		}

		http.Error(w, msg, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		msg := fmt.Sprintf("%s %s successfully added to Rolodex.", contact.FirstName, contact.LastName)
		index.ExecuteTemplate(w, "layout", msg)
	}
}
