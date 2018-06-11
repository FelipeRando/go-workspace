/*
Take the previous program in the previous folder and change it so that:
a template is parsed and served
you pass data into the template
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	data := 42
	err := t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "au au")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Random")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
