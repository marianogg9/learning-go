package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	meters := feet * 0.3048

	if err != nil {
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
