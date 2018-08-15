package commands

import (
	"flag"
)

func SaySomething() *string {
	returnString := flag.String("say-something", "nada escrito", "type anything to say")
	flag.Parse()
	return returnString
}
