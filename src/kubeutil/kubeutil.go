package main

import (
	"fmt"
	"kubeutil/commands"
)

func main() {
	saySomething := *commands.SaySomething()
	fmt.Println(saySomething)
}
