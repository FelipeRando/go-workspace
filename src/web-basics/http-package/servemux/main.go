package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type avocado int

//THIS IS NOT A HANDLER!!!
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index de uma função")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "everything is fine, I'm listening for requests ;)")
}

//THIS IS A HANDLER!!!
func (avd avocado) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "brought you by a handler")
}

func main() {
	var av avocado
	mux := http.NewServeMux()
	//handlefunc != handle
	mux.HandleFunc("/", index)
	mux.Handle("/handler", av)
	//handlefunc passa a nossa função para uma outra função chamanda Handle
	mux.HandleFunc("/teste", healthCheck)
	fmt.Println("starting server...")
	go testServerConnection()
	http.ListenAndServe(":8080", mux)
}

func testServerConnection() {
	time.Sleep(5)
	r, err := http.Get("http://localhost:8080/teste")
	if err != nil {
		log.Fatal(err)
		return
	}
	bs := make([]byte, 99999)
	r.Body.Read(bs)
	fmt.Println(string(bs))
}
