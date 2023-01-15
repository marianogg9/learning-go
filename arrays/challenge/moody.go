package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main1() {
	if len(os.Args) != 2 {
		fmt.Printf("Please give me a name\n")
		return
	} else {
		moods := [...]string{
			"terrible X(",
			"sad :(",
			"happy :)",
			"awesome :D",
		}
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(moods))

		fmt.Println(os.Args[1] + " feels " + moods[n])
	}
}
