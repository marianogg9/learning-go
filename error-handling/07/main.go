package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	year := os.Args[1]

	if y, err := strconv.Atoi(year); err != nil {
		fmt.Printf("%v is not a valid year\n", year)
		return
	} else if y%4 == 0 {
		fmt.Printf("%v is a leap year.\n", y)
	} else {
		fmt.Printf("%v is not a leap year.\n", y)
	}
}
