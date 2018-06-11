/*
Take the previous program and change it so that:
func main uses http.Handle instead of http.HandleFunc
Contstraint: Do not change anything outside of func main
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
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
