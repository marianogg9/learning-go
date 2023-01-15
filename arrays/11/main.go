package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	nums := [5]float64{}
	var sum float64

	for i := range args {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		nums[i] = n
		sum += n
	}

	fmt.Printf("Your numbers: %v\n", nums)
	fmt.Println("Average: ", sum/5)
}
