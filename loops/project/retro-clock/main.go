package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"|||",
		"| |",
		"| |",
		"| |",
		"|||",
	}

	one := placeholder{
		" ||",
		"  |",
		"  |",
		"  |",
		"  |",
	}

	two := placeholder{
		"|||",
		"  |",
		"|||",
		"|  ",
		"|||",
	}

	three := placeholder{
		"|||",
		"  |",
		"|||",
		"  |",
		"|||",
	}

	four := placeholder{
		"| |",
		"| |",
		"|||",
		"  |",
		"  |",
	}

	five := placeholder{
		"|||",
		"|  ",
		"|||",
		"  |",
		"|||",
	}

	six := placeholder{
		"|||",
		"|  ",
		"|||",
		"| |",
		"|||",
	}

	seven := placeholder{
		"|||",
		"  |",
		"  |",
		"  |",
		"  |",
	}

	eight := placeholder{
		"|||",
		"| |",
		"|||",
		"| |",
		"|||",
	}

	nine := placeholder{
		"|||",
		"| |",
		"|||",
		"  |",
		"|||",
	}

	colon := placeholder{
		"   ",
		" : ",
		"   ",
		" : ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10], // this works as type(hour) = int, when integer divided by 10 it will result in either 0, 1 or 2 which is then the index of the 'digits' slice to be printed.
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], " ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
