package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("[command] [string]\n\nAvailable commands: lower, upper and title\n")
		return
	}

	c, s, o := os.Args[1], os.Args[2], ""

	switch c {
	case "lower":
		o = strings.ToLower(s)
	case "upper":
		o = strings.ToUpper(s)
	case "title":
		o = strings.Title(s)
	default:
		fmt.Printf("Unknown command %v", c)
		return
	}

	fmt.Printf("%s\n", o)
}
