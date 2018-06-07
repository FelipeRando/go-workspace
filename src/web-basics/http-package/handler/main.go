package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var t *template.Template

type handler int

func init() {
	t = template.Must(t.ParseGlob("*.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Method       string
		URL          *url.URL
		Submissions  map[string][]string
		Header       http.Header
		Host         string
		ContentLengh int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	err = t.Execute(w, data)
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
