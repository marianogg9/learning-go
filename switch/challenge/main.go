package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Now: %v\n", t)
	switch {
	case t.Hour() > 18:
		fmt.Printf("Good evening\n")
	case t.Hour() > 12:
		fmt.Printf("Good afternoon\n")
	case t.Hour() >= 6:
		fmt.Printf("Good morning\n")
	default:
		fmt.Printf("Good night\n")
	}
}
