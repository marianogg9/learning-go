package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxTurns = 5
	win      = "Yes! Congrats! Kudos!"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Pick a num.\n")
		return
	}

	looseMessages := [3]string{"You lost!", "What a pitty!", "Try again!"}

	rand.Seed(time.Now().UnixNano())

	guess, err := strconv.Atoi(os.Args[1])

	winMessages := strings.Fields(win)

	if err != nil {
		fmt.Printf("Not a num.\n")
		return
	}

	if guess < 0 {
		fmt.Printf("Pick a num > 0.\n")
		return
	}

	for t := 0; t < maxTurns; t++ {
		n := rand.Intn(guess + 1)
		if t == 0 {
			if n == guess {
				fmt.Printf("%v In the first attempt!\n", winMessages[rand.Intn(len(winMessages))])
				return
			}
		} else {
			if n == guess {
				fmt.Printf("%v\n", looseMessages[rand.Intn(len(looseMessages))])
				return
			}
		}
		//fmt.Println(n)
	}
	println("You lost!")
}
