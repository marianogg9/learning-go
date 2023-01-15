package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	const (
		EUR = iota
		GBP
		JPY
	)

	ratios := [...]float64{
		EUR: 0.99,
		GBP: 0.87,
		JPY: 145.26,
	}

	usd, err := strconv.ParseFloat(args[0], 64)
	if err == nil {
		fmt.Printf("%.2f USD is %.2f EUR\n", usd, (ratios[EUR] * usd))
		fmt.Printf("%.2f USD is %.2f GBP\n", usd, (ratios[GBP] * usd))
		fmt.Printf("%.2f USD is %.2f JPY\n", usd, (ratios[JPY] * usd))
	} else {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

}
