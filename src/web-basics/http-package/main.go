package main

import (
	"fmt"
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
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, r.PostForm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%+v", r)
	fmt.Println(r.PostForm)
}

func main() {
	fmt.Println("Starting server...")
	var h handler
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatal(err)
	}

}
