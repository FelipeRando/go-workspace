//cookies are a way to store information about clients, a way to create state
//if the client refuses to use cookies we can also use unique IDs on each request to identify this user
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	pointerToCookie := &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	}
	http.SetCookie(w, pointerToCookie)
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
