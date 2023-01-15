package main

import (
	"fmt"
	"path"
)

func main() {
	var (
		hola string
	)

	hola, _ = path.Split("secret/file.txt")

	fmt.Println(hola)
}
