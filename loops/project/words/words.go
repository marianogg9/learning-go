// given a list of words and a query, return the search output.

package main

import (
	"fmt"
	"os"
	"strings"
)

const campus = "lazy cat jumps again and again and again"

func word() {
	words := strings.Fields(campus)
	query := os.Args[1:] // to get all args except 1st (program OS path)

	for _, q := range query {
		for i, w := range words {
			// if strings.ToLower(q) == strings.ToLower(w) { // case insensitive search
			if strings.EqualFold(w, q) { // fancier!!
				fmt.Printf("#%-2d: %q\n", i+1, w) // i+1 to output position instead of index

				break
			}
		}
	}
}

func wordBreakLabel() {
	words := strings.Fields(campus)
	query := os.Args[1:]

queries: // break label
	for _, q := range query {
		for i, w := range words {
			if strings.EqualFold(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				break queries // it will break parent loop
			}
		}
	}
}

func wordContinueLabel() {
	words := strings.Fields(campus)
	query := os.Args[1:]

queries: // continue label
	for _, q := range query {
		for i, w := range words {
			if strings.EqualFold(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				continue queries // it will continue from parent loop
			}
		}
	}
}

func wordSwitchLabel() {
	words := strings.Fields(campus)
	query := os.Args[1:]

queries:
	for _, q := range query {
	search:
		for i, w := range words {
			switch strings.ToLower(q) {
			case "and", "or", "the":
				break search // breaks search loop (inner) to skip filtered words
			}
			if strings.EqualFold(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				continue queries
			}
		}
	}
}

func main() {
	//word()
	//wordBreakLabel()
	//wordContinueLabel()
	//wordBreakLabel()
	wordSwitchLabel()
}
