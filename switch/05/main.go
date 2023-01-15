package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	m := strings.ToLower(month)

	switch m {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		switch leap {
		default:
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	// if m := strings.ToLower(month); m == "april" ||
	// 	m == "june" ||
	// 	m == "september" ||
	// 	m == "november" {
	// 	days = 30
	// } else if m == "january" ||
	// 	m == "march" ||
	// 	m == "may" ||
	// 	m == "july" ||
	// 	m == "august" ||
	// 	m == "october" ||
	// 	m == "december" {
	// 	days = 31
	// } else if m == "february" {
	// 	if leap {
	// 		days = 29
	// 	}
	// } else {
	// 	fmt.Printf("%q is not a month.\n", month)
	// 	return
	// }

	fmt.Printf("%q has %d days.\n", month, days)
}
