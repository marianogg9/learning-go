package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		min, errmi := strconv.Atoi(os.Args[1])
		max, errma := strconv.Atoi(os.Args[2])

		if errmi == nil && errma == nil && max > min {
			s := 0
			for i := min; i < max; i++ {
				fmt.Printf("%v + ", i)
				s += i
			}
			s += max
			fmt.Printf("%v = %v\n", max, s)
		} else {
			fmt.Printf("Arguments not valid.\n")
			return
		}
	} else {
		fmt.Printf("Usage: run main.go [min] [max]\n")
		return
	}
}
