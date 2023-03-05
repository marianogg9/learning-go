package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	var items []byte

	sort.Strings(args)

	for ind, i := range args {
		item := strconv.Itoa(ind+1) + ". " + i // add ordinal to each sorted element
		items = append(items, item...)
		items = append(items, '\n')
	}

	err := ioutil.WriteFile("slices/2-sort-to-a-file/sorted.txt", items, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
