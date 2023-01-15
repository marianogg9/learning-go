package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	var nums []int

	datal := strings.Split(data, "")

	for i := 0; i < len(datal); i++ {
		if j, err := strconv.Atoi(datal[i]); err == nil {
			nums = append(nums, j)
		}
	}

	var (
		numsEven, numsOdd []int
	)
	for _, i := range nums {
		if i%2 == 0 {
			numsEven = append(numsEven, i)
		} else {
			numsOdd = append(numsOdd, i)
		}
	}
	fmt.Println("nums:", nums)
	fmt.Println("evens:", numsEven)
	fmt.Println("odds:", numsOdd)
	fmt.Println("middle:", nums[2:4])
	fmt.Println("first 2:", nums[0:2])
	fmt.Println("last 2:", nums[len(nums)-2:])
	fmt.Println("evens last 1:", numsEven[len(numsEven)-1:])
	fmt.Println("odds last 2:", numsOdd[len(numsOdd)-2:])
}
