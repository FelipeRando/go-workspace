package main

import (
	"flag"
	"fmt"
)

func main() {
	sayHello := flag.String("say-hello", "Thanks for using kubeutil", "this command says hello")
	flag.Parse()
	fmt.Println(*sayHello)
}
