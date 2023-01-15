package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	l, err := strconv.Atoi(os.Args[1])

	if err == nil {

		fmt.Printf("%5s", "X")

		for i := 0; i <= l; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 0; i <= l; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= l; j++ {
				fmt.Printf("%5d", i*j)
			}
			fmt.Println()
		}
	}
}
