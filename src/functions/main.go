package main 

import (
	"fmt"
	"functions/mat"
)

func main() {
	soma := matematica.Add(2,2)
	multiplicacao := matematica.Multiply(3,2)
	fmt.Println(soma)
	fmt.Println(multiplicacao)

	//Closure: function inside a function
	count := 0
	addToCount := func() int{
		count ++
		return count
	}
	fmt.Println(addToCount())
	fmt.Println(addToCount())

	even := evenGenerator()
	fmt.Println(even())
	fmt.Println(even())

	fatorial := factorial(99)
	fmt.Println(fatorial)
}

//Closure
func evenGenerator() func() int{
	i := 0
	return func() int{
		i += 2
		return i
	}
}

//Recursion
func factorial(x int) int {
	if x == 0{
		return 1
	}
	return x * factorial(x-1)
}

