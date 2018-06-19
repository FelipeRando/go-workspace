package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", formRequest)
	http.ListenAndServe(":8080", nil)
}

func formRequest(w http.ResponseWriter, r *http.Request) {

}
