package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Println("There are %d names.\n", count)
}
