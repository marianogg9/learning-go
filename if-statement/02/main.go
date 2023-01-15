package main

import "fmt"

func main() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere")
	} else {
		fmt.Println("I don't know")
	}

	// if radius >= 200 {
	// 	big = true
	// } else {
	// 	big = false
	// }

	// if isSphere && radius >= 200 {
	// 	fmt.Println("It's a big sphere")
	// } else {
	// 	fmt.Println("I don't know")
	// }

	// if big != true {
	// 	fmt.Println("I don't know.")
	// } else if !(isSphere == false) {
	// 	fmt.Println("It's a big sphere.")
	// } else {
	// 	fmt.Println("I don't know.")
	// }
}
