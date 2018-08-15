package main

import (
	"flag"
	"fmt"
	"kubeutil/commands"
)

func main() {
	saySomething := *commands.SaySomething()
	flag.Parse()
	fmt.Println(saySomething)

}
