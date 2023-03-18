package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
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

	for md := 0; md < 3; md++ {
		for mc := 0; mc < 10; mc++ {
			if md == 2 && mc == 4 {
				break
			}
			for d := 0; d < 6; d++ { // decenas d<6
				for c := 0; c < 3; c++ { // solo imprime el numero c<10
					screen.Clear()
					screen.MoveTopLeft()
					for ml := range digits[mc] {
						fmt.Printf("%v ", digits[md][ml])
						fmt.Printf("%-5v\n", digits[mc][ml])
					}
					for l := range digits[c] {
						fmt.Printf("%15v ", digits[d][l])
						fmt.Printf("%v\n", digits[c][l])
					}
					time.Sleep(1 * time.Second / 20)
				}
			}
		}
	}

}
