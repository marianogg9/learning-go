package main

import "fmt"

func main() {
	s := 0
	for i := 1; i <= 10; i++ {
		s += i
		if i < 10 {
			fmt.Printf("%d + ", i)
		} else {
			fmt.Printf("%d = %d\n", i, s)
		}
	}
}
