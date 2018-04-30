package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//we have to create byte slice using make and n empty spaces
	//because the Read function is not configured to resize the slice if it is full
	r := make([]byte, 9999)
	resp.Body.Read(r)
	fmt.Println(string(r))

	lw := logWriter{}

	//the same thing is here
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
