package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		ptoppings []string
		dtimes    []time.Time
		gyears    []int
		lights    []bool
	)

	now := time.Now()

	ptoppings = append(ptoppings, "huevito", "mayonesa", "aceitunitas")
	dtimes = append(dtimes, now.Add(time.Hour*24), now.Add(time.Minute*54))
	gyears = append(gyears, 1988, 2019, 2022)
	lights = append(lights, true, false)

	fmt.Println(ptoppings)
	fmt.Println(dtimes)
	fmt.Println(gyears)
	fmt.Println(lights)
}
