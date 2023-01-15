package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesASplit := strings.Split(namesA, ", ")
	slices.Sort(namesASplit)
	slices.Sort(namesB)

	equal := true
	if len(namesB) == len(namesASplit) {
		for i, v := range namesASplit {
			if v != namesB[i] {
				equal = false
			}
		}
	} else {
		equal = false
	}

	fmt.Println(equal)
}
