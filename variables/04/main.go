package main

import "fmt"

func main() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2*width + 2*height

	fmt.Println(perimeter)
}
