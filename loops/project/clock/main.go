package main

import (
	"fmt"
	"time"
)

func main() {

	digits := [10][5]string{
		{"|||", "| |", "| |", "| |", "|||"},
		{" ||", "  |", "  |", "  |", "  |"},
		{"|||", "  |", "|||", "|  ", "|||"},
		{"|||", "  |", "|||", "  |", "|||"},
		{"| |", "| |", "|||", "  |", "  |"},
		{"|||", "|  ", "|||", "  |", "|||"},
		{"|||", "|  ", "|||", "| |", "|||"},
		{"|||", "  |", "  |", "  |", "  |"},
		{"|||", "| |", "|||", "| |", "|||"},
		{"|||", "| |", "|||", "  |", "  |"},
	}

	for md := 0; md < 6; md++ {
		for mc := 0; mc < 10; mc++ {
			for d := 0; d < 2; d++ { // decenas d<6
				for c := 0; c < 3; c++ { // solo imprime el numero c<10
					for ml := range digits[mc] {
						fmt.Printf("%v ", digits[md][ml])
						fmt.Printf("%-5v\n", digits[mc][ml])
					}
					for l := range digits[c] {
						fmt.Printf("%v ", digits[d][l])
						fmt.Printf("%-5v\n", digits[c][l])
					}
					time.Sleep(1 * time.Second)
					fmt.Println()
				}
			}
		}
	}

}
