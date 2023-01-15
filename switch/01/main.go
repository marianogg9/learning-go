package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	e, err := strconv.ParseFloat(os.Args[1], 32)

	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch err == nil {
	case e < 2.0:
		fmt.Printf("%v is micro\n", e)
	case e >= 2.0 && e < 3.0:
		fmt.Printf("%v is very minor\n", e)
	case e >= 3.0 && e < 4.0:
		fmt.Printf("%v is minor\n", e)
	case e >= 4.0 && e < 5.0:
		fmt.Printf("%v is light\n", e)
	case e >= 5.0 && e < 6.0:
		fmt.Printf("%v moderate\n", e)
	case e >= 6.0 && e < 7.0:
		fmt.Printf("%v is strong\n", e)
	case e >= 7.0 && e < 8.0:
		fmt.Printf("%v is major\n", e)
	case e >= 8.0 && e < 9.0:
		fmt.Printf("%v is great\n", e)
	default:
		fmt.Printf("%v is massive\n", e)
	}
}
