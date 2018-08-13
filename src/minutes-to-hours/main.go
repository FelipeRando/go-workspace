package main

import "fmt"

// MinsToHours converts Minutes to Hours and Minutes
func MinsToHours(mins int) string {
	return fmt.Sprintf("%d:%.2d", mins/60, mins%60)
}

func main() {
	fmt.Println(MinsToHours(500))
}
