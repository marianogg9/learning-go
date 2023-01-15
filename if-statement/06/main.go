package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
	user     = "jack"
	pass     = "1888"
	user1    = "inanc"
	pass1    = "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user {
		if u == user1 && p == pass1 {
			fmt.Printf(accessOK, u)
		} else {
			fmt.Printf(errUser, u)
		}
	} else if p != pass {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
