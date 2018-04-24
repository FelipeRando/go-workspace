package main

import "fmt"

func main() {
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, number := range numbers {
		if number %2 == 1 {
			fmt.Printf("Number %v is odd\n", number)
		} else {
			fmt.Printf("Number %v is even\n", number)
		}
	}
}
