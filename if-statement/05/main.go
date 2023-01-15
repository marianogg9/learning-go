package main

import (
	"fmt"
	"os"
)

func main() {
	username := "jack"
	password := "1888"

	l := os.Args

	if len(l)-1 < 2 {
		fmt.Printf("Usage: [username] [password]\n")
	} else if os.Args[1] != username {
		fmt.Printf(`Access denied for "` + os.Args[1] + `"` + "\n")
	} else if os.Args[2] != password {
		fmt.Printf(`Access denied for "jack".` + "\n")
	} else {
		fmt.Printf(`Access granted for "jack".` + "\n")
	}
}
