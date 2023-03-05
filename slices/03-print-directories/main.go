package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}
	for _, i := range args {
		var dir []byte
		fmt.Printf("input directory: %v\n", i)
		dirs, err := ioutil.ReadDir(i) // get all files in input dir
		if err != nil {                // error handling
			fmt.Println(err)
			return
		}
		for _, d := range dirs {
			if os.FileInfo.IsDir(d) { // if the item is a directory, append its name and repeat
				//fmt.Printf("%v/\n", d.Name())
				dir = append(dir, d.Name()...)
				dir = append(dir, '/')
				dir = append(dir, '\n')
				dir = append(dir, '\t')
				d, err := ioutil.ReadDir(d.Name())
				if err != nil { // error handling
					fmt.Println(err)
					continue
				}
				var subdir []byte
				for _, sdir := range d {
					if os.FileInfo.IsDir(sdir) { // append all current directory content (that are directories themselves)
						//fmt.Printf("\t%v/\n", sdir.Name())
						subdir = append(subdir, sdir.Name()...)
						subdir = append(subdir, '/')
						subdir = append(subdir, '\n')
						subdir = append(subdir, '\t')
					}
				}
				dir = append(dir, subdir...)
				dir = append(dir, '\n')
			}
		}
		e := ioutil.WriteFile("slices/03-print-directories/dirs.txt", dir, 0644)
		if e != nil { // error handling
			fmt.Println(e)
			return
		}
	}

}
