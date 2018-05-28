package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	cs := make(chan string) //create a channel

	for _, link := range links {
		go checkLink(link, cs)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-cs)
	}
}

func checkLink(link string, cs chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		cs <- "Might be down"
		return
	}

	fmt.Println(link, "is up")
	cs <- "Yep its up"
}
