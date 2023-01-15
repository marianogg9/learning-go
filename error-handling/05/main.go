package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := os.Args[1]

	if n, err := strconv.Atoi(age); len(os.Args) < 2 || err != nil {
		fmt.Printf("Requires age\n")
	} else if n > 17 {
		fmt.Printf("R-Rated\n")
	} else if n <= 17 && n >= 13 {
		fmt.Printf("PG-13\n")
	} else if n < 13 && n > 0 {
		fmt.Printf("PG-Rated\n")
	} else {
		fmt.Printf("Wrong age: %v\n", n)
	}
}
