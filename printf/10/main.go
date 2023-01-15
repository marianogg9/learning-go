package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Your mame is %v and your lastname is %v\n", os.Args[1], os.Args[2])
}
