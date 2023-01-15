package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	query := os.Args
	if len(query) != 2 {
		fmt.Println("Usage: go run main.go [query]\n")
		return
	} else {
		path := os.Getenv("PATH")
		path_list := filepath.SplitList(path)
		for i, pl := range path_list {
			if query[1] == pl {
				fmt.Printf("#%-2d: %q\n", i, pl)
				return
			}
		}
		fmt.Printf("%v not found in $PATH\n", query[1])
	}
}
