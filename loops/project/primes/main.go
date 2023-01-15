package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	q := os.Args[1]
	i, err := strconv.ParseUint(q, 10)
	if err == nil && isPrime(i) {
		fmt.Printf("%v is prime\n", q)
	} else {
		fmt.Printf("is not\n")
	}
}

func isPrime(n int) bool {
	//var prime bool

	if n == 2 {
		return true
	}
	if n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}

	i := 5
	w := 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}
	return true
}
