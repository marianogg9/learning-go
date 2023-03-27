package main

import "fmt"

func main() {
	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, i := range words {
		w := []byte(i)
		fmt.Println(w)
		bwords = append(bwords, w)
	}

	for _, bw := range bwords {
		fmt.Println(string(bw))
	}
}
