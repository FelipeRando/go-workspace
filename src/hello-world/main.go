package main

import "fmt"
import "hello-world/hello"

func main() {
	fmt.Println("Hello World!")
	println("esse foi printado usando println built in")

	a := 1
	b := 2

	resultado := maior(a, b)
	if resultado == false {
		fmt.Println("falso")
	}
	//using your own package
	hello.SayHello()
}

// function that returns true or false
func maior(a, b int) bool {
	return a > b
}
