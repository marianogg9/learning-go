package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	m := os.Args[1]

	switch m {
	case "micro":
		fmt.Printf("%v's richter scale is less than 2.0\n", m)
	case "very minor":
		fmt.Printf("%v's richter scale is 2 - 2.9\n", m)
	case "minor":
		fmt.Printf("%v's richter scale is 3 - 3.9\n", m)
	case "light":
		fmt.Printf("%v's richter scale is 4 - 4.9\n", m)
	case "moderate":
		fmt.Printf("%v's richter scale is 5 - 5.9\n", m)
	case "strong":
		fmt.Printf("%v's richter scale is 6 - 6.9\n", m)
	case "major":
		fmt.Printf("%v's richter scale is 7 - 7.9\n", m)
	case "great":
		fmt.Printf("%v's richter scale is 8 - 8.9\n", m)
	case "massive":
		fmt.Printf("%v's richter scale is 10+\n", m)
	default:
		fmt.Printf("%v's richter scale is unknown\n", m)
	}
}
