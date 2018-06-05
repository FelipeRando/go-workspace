package main

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

type handler int

func init() {
	t = template.Must(t.ParseGlob("*.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, 42)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var h handler
	http.ListenAndServe(":8080", h)
}
