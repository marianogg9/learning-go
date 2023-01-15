package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)
}

func main1() {
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}

func main2() {
	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("SUCESS: Converted %q to %d.\n", age, n)
}
