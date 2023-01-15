package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		errorMsg  = "Provide only the [starting] and [stopping] positions."
		errorConv = "The provided parameters are not integers."
		errorPos  = "Wrong positions."
	)
	var (
		from, to int
		output   []string
	)
	// uncomment the slice below
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	args := os.Args[1:]

	if len(args) == 2 {
		if i, err := strconv.Atoi(args[0]); err == nil {
			from = i
		} else {
			fmt.Println(errorConv)
			return
		}
		if i, err := strconv.Atoi(args[1]); err == nil {
			to = i
		} else {
			fmt.Println(errorConv)
			return
		}
	} else if len(args) == 1 {
		if i, err := strconv.Atoi(args[0]); err == nil {
			from = i
			to = len(ships)
		}
	} else {
		fmt.Println(errorMsg)
		return
	}

	if to <= len(ships) {
		if from <= to && from >= 0 {
			output = ships[from:to]
			fmt.Println(output)
		} else {
			fmt.Println(errorPos)
			return
		}
	}
}
