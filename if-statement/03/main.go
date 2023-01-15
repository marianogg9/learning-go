package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args)
	if count < 2 {
		fmt.Println("Give me args")
	} else if count == 2 {
		fmt.Printf("There is one: \"%s\"\n", os.Args[1])
	} else if count == 3 {
		fmt.Printf("There are two: \"%s %s\"\n", os.Args[1], os.Args[2])
	} else {
		fmt.Println("There are", count, "arguments")
	}
}
