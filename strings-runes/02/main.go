package main

import (
	"fmt"
	"strings"
)

func main() {
	const word = "console"

	var s []byte

	fmt.Printf("%-5v %-5v %-5v\n", "dec", "hexa", "binary")
	fmt.Printf("%v\n", strings.Repeat("-", 45))
	for _, l := range word {
		fmt.Printf("%-5d %-5x %-5b\n", l, l, l)
		// s = append(s, []byte(string(l)))
		// sd = append(sd, word[i])
	}
	fmt.Println(s)

	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
