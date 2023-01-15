package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		c := os.Args[1]

		if c != "*" || c != "/" || c != "+" || c != "-" {
			fmt.Printf("Invalid operator.\nUsage: [op=*/+-] [size]\n")
			return
		}

		l, err := strconv.Atoi(os.Args[2])

		if err == nil && l > 0 {
			fmt.Printf("%5v", c)

			for i := 0; i <= l; i++ {
				fmt.Printf("%5v", i)
			}
			fmt.Println()

			for i := 0; i <= l; i++ {
				fmt.Printf("%5v", i)

				for j := 0; j <= l; j++ {
					var op int

					switch c {
					case "*":
						op = i * j
					case "/":
						if j != 0 {
							op = i / j
						} else {
							op = 0
						}
					case "+":
						op = i + j
					case "%":
						if j != 0 {
							op = i % j
						} else {
							op = 0
						}
					default:
						op = i - j
					}

					fmt.Printf("%5v", op)
				}
				fmt.Println()
			}
			fmt.Println()
		} else {
			fmt.Println("Invalid length of the table.")
			return
		}

	} else {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}
}
