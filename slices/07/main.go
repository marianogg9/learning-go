package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice
	header = append(header, 'P', 'N', 'G')

	// 3. compare the slices using the bytes.Equal function
	if !bytes.Equal(png, header) {
		return
	} else {
		fmt.Println("They are equal.")
	}

	// 4. print whether they're equal or not
}
