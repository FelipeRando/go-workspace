package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	defer fmt.Println("Tchau amiguinho")
	http.HandleFunc("/doggy", doggy)
	http.HandleFunc("/pipboy", fallout)
	http.Handle("/", http.FileServer(http.Dir("."))) //FilServer is the most common way to serve files using go
	http.ListenAndServe(":8080", nil)
}

//serve with io.Copy
func doggy(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpeg")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

//serve with http.ServeFile
func fallout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pipboy.jpg")
}
