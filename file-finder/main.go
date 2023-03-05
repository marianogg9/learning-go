package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a dir name")
		return
	}
	files, err := ioutil.ReadDir(args[0]) // get all file(s) info from an input directory
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for _, f := range files { // calculate how much memory to reserve in backing array
		if f.Size() == 0 { // get all empty files (size == 0)
			total += len(f.Name()) + 1
		}
	}
	fmt.Println("Total required space: %d bytes.\n", total)

	//var names []byte
	names := make([]byte, 0, total) // create final slice with exact allocated space

	for _, f := range files {
		if f.Size() == 0 { // get all empty files (size == 0)
			name := f.Name()
			fmt.Println(name) // print their names
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	err = ioutil.WriteFile("file-finder/out.txt", names, 0644) // write found empty files to a file
	if err != nil {
		fmt.Println(err)
		return
	}

}
