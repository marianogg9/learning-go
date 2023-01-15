package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("provide a few numbers")
		return
	} else {
		var nums []int
		for _, i := range args {
			if n, err := strconv.Atoi(i); err == nil {
				nums = append(nums, n)
			}
		}
		sort.Ints(nums)
		fmt.Println(nums)
	}
}
