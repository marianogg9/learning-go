package main

import "fmt"

func main() {
	const (
		Winter = 12 - 3*iota
		Fall   = 9
		Summer = 6
		Spring = 3
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
