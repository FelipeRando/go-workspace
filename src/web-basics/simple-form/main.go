package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type person struct {
	Name string
}

func main() {
	http.HandleFunc("/", printForm)
	http.ListenAndServe(":8080", nil)
}

func printForm(w http.ResponseWriter, r *http.Request) {
	pn := r.FormValue("name")
	err := tpl.Execute(w, person{pn})
	if err != nil {
		log.Fatal(err)
	}
}
