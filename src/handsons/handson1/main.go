package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	lado float64
}

type circle struct {
	raio float64
}

func (c circle) getArea() float64 {
	return 3.14 * c.raio
}

func (s square) getArea() float64 {
	return s.lado * s.lado
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{4}
	c := circle{2}

	printArea(s)
	printArea(c)
}
