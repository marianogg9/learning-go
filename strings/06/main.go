package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	fmt.Println(strings.ToLower(msg))
}
