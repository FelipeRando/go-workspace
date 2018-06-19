//cookies are a way to store information about clients, a way to create state
//if the client refuses to use cookies we can also use unique IDs on each request to identify this user
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	pointerToCookie := &http.Cookie{
		Name:  "my-cookie",
		Value: "0",
	}
	//you can pass multiple cookies in one request, just add cookies in another call to SetCookie function
	count, _ := strconv.Atoi(pointerToCookie.Value) //convert string to int
	count++
	pointerToCookie.Value = strconv.Itoa(count) //convert int to string
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
