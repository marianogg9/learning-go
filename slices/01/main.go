package main

import (
	"fmt"
)

func main() {
	var names []string // declaration as nil
	var distances []int
	var data []uint8
	var ratios []float64
	var alives []bool

	names = []string{} // assign empty

	fmt.Printf("%T %v %v\n", names, len(names), names == nil)
	fmt.Printf("%T %v %v\n", distances, len(distances), distances == nil)
	fmt.Printf("%T %v %v\n", data, len(data), data == nil)
	fmt.Printf("%T %v %v\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("%T %v %v\n", alives, len(alives), alives == nil)
}
