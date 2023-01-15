package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	l := len(os.Args)
	// if l != 2 {
	// 	fmt.Println("Give me a letter")
	// } else if os.Args[1] == "a" || os.Args[1] == "e" || os.Args[1] == "i" || os.Args[1] == "o" || os.Args[1] == "u" {
	// 	fmt.Printf("\"" + os.Args[1] + "\"" + " is a vowel.\n")
	// } else if os.Args[1] == "y" || os.Args[1] == "w" {
	// 	fmt.Printf("\"" + os.Args[1] + "\"" + " is sometimes a vowel, sometimes not.\n")
	// } else {
	// 	fmt.Printf("\"" + os.Args[1] + "\"" + " is a consonant.\n")
	// }

	if l != 2 {
		fmt.Println("Give me a letter")
	} else if strings.ContainsAny(os.Args[1], "aeiou") {
		fmt.Printf("\"" + os.Args[1] + "\"" + " is a vowel.\n")
	} else if strings.ContainsAny(os.Args[1], "yw") {
		fmt.Printf("\"" + os.Args[1] + "\"" + " is sometimes a vowel, sometimes not.\n")
	} else {
		fmt.Printf("\"" + os.Args[1] + "\"" + " is a consonant.\n")
	}
}
