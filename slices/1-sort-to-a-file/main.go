package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	var items []byte

	sort.Strings(args)

	for _, i := range args {
		items = append(items, i...)
		items = append(items, '\n')
	}

	err := ioutil.WriteFile("slices/1-sort-to-a-file/sorted.txt", items, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
