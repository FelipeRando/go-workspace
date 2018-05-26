package main

import ("fmt")

type shape interface{
	getArea() float64
}

type triangle struct{
	height float64
	base float64
	name string
}

type square struct{
	sideLenght float64
	name string
}

func main(){
	s := square{4,"square"}
	t := triangle{4,3,"triangle"}

	printArea(s)
	printArea(t)
}

func (t triangle) getArea() float64{
	return (0.5 * t.base * t.height)
} 

func (s square) getArea() float64{
	return s.sideLenght * s.sideLenght
}

func printArea(s shape){
	fmt.Println("The area of","is",s.getArea())
}