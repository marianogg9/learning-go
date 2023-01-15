package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "inanç           "

	fmt.Println(utf8.RuneCountInString(strings.TrimRight(name, " ")))
}
