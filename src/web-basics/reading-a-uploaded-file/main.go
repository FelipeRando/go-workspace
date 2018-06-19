package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", readFromFile)
	http.ListenAndServe(":8080", nil)
}

func readFromFile(w http.ResponseWriter, r *http.Request) {

	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {

		//open the file
		f, h, err := r.FormFile("q") //q is the name of the input on html page
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close() //this file must be closed sometime

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		//read from the file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
		</form>
		<br>
		`+s)
}
