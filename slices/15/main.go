package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
				New York,100,2,1,100000
				New York,150,3,2,200000
				Paris,200,4,3,400000
				Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		cities   []string
		sizes    []int
		beds     []int
		baths    []int
		prices   []int
		from, to int
	)

	lines := strings.Split(data, "\n")
	//fmt.Println("lines:", lines)
	//var columns []string
	// for _, i := range lines {
	// 	fmt.Println(i)
	// 	columns = append(columns)
	// }

	for i := 1; i < len(lines); i++ {
		column := strings.Split(lines[i], separator)
		//fmt.Println("column:", column)
		cities = append(cities, strings.Trim(column[0], "\t "))
		if s, err := strconv.Atoi(column[1]); err == nil {
			sizes = append(sizes, s)
		}
		if b, err := strconv.Atoi(column[2]); err == nil {
			beds = append(beds, b)
		}
		if ba, err := strconv.Atoi(column[3]); err == nil {
			baths = append(baths, ba)
		}
		if p, err := strconv.Atoi(column[4]); err == nil {
			prices = append(prices, p)
		}
	}
	// fmt.Println(cities)
	// fmt.Println(sizes)
	// fmt.Println(beds)
	// fmt.Println(baths)
	// fmt.Println(prices)

	args := os.Args[1:]
	//fmt.Println(args)
	to = len(strings.Split(lines[0], separator)) - 1
	from = 0
	fmt.Println("lines[0]:", lines[0])
	switch len(args) {
	case 2:
		for i, c := range strings.Split(lines[0], separator) {
			if c == args[0] {
				from = i
			} else {
				if c == args[1] {
					to = i
				}
			}
		}
		fmt.Println("to:", to)
		fmt.Println("from:", from)
	case 1:
		for i, c := range strings.Split(lines[0], separator) {
			if c == args[0] {
				from = i
			}
		}
		fmt.Println("to:", to)
		fmt.Println("from:", from)
	case 0:
		fmt.Println("to:", to)
		fmt.Println("from:", from)
	default:
		fmt.Println("Wrong number of parameters.")
		return
	}
}
