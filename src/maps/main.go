package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "46bc00",
	}
	colors["white"] = "fffffff"
	delete(colors, "red")

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Color: %v have this hex ---> %v\n", color, hex)
	}
}
