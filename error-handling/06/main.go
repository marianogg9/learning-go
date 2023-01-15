package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Pick a number\n")
		return
	}

	num := os.Args[1]

	if n, err := strconv.Atoi(num); n%8 == 0 && err == nil {
		fmt.Printf("%v is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 && err == nil {
		fmt.Printf("%v is an even number\n", n)
	} else if err == nil {
		fmt.Printf("%v is an odd number\n", n)
	} else {
		fmt.Printf("%v is not a number\n", n)
	}
}
