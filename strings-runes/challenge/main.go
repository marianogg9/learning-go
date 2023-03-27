package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Too much or not enough args, bye")
		return
	}

	ar := strings.Join(args, "")
	buf := make([]byte, len(ar))

	// la condicion para que sea url es que arranque con //

	anterior := ar[0]
	for i := 1; i < len(ar); i++ {
		// fmt.Println(strings.Repeat("-", 5))
		// fmt.Println("current char: ", string(ar[i]))
		// fmt.Println("current outside index: ", i)
		if ar[i] == 47 { // '/'
			buf = append(buf, ar[i])

			if anterior == 47 { // '//'
				for indice, x := range ar[i+1:] {
					//fmt.Println("current inner char:", string(x))
					if x != ' ' {
						//fmt.Println("indice: ", indice)
						buf = append(buf, '*')
					} else {
						//fmt.Println(x)
						i += indice
						break
					}
				}
			}
		} else {
			buf = append(buf, ar[i])
		}
		anterior = ar[i]
		//fmt.Println(string(buf))
	}

	fmt.Println(string(buf))

}
