package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// In Go, rune = Unicode Code Point.
// A rune literal is a typeless number, so it can represent a byte, rune, or any other type.

func main() {
	// num := []byte{'h', 'e', 'y'}
	// fmt.Println(string(num))

	var start, stop int
	if args := os.Args[1:]; len(args) == 2 { // get start/stop from parameters
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 { // if incorrect parameters, assign default values
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-10s\n%s\n", "literal", "dec", "hex", "encoded", strings.Repeat("-", 45))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -10x\n", n, string(n)) // c prints a char by the given code point, d is decimal, x hexa
	}
}
