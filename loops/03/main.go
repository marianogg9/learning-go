package main

import "fmt"

func main() {
	s := 0
	for i := 0; i <= 10; i++ {
		s += i
	}
	fmt.Printf("Sum: %d\n", s)
}
