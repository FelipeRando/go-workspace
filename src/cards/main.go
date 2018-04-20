package main

import (
	"fmt"
)

func main() {
	var slice []int
	slice = append(slice, 1)
	fmt.Println(slice)

	slice2 := []string{"olkoko"}
	slice2 = append(slice2, "2")
	fmt.Println(slice2)

	for i, value := range slice2 {
		i++
		fmt.Printf("%dst value of slice is %s\n", i, value)
	}
}
