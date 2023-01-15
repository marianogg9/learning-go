package main

import (
	"fmt"
	"os"
)

func main() {
	a, b, c := os.Args[1], os.Args[2], os.Args[3]
	fmt.Println("There are %s people!", len(os.Args)-1)
	fmt.Println("Hello great", a, "!\nHello great", b, "!\nHello great", c, "!")
	fmt.Println("Nice to meet you all.")
}
