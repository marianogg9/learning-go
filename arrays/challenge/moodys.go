package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("[name] [positive|negative]")
		return
	} else {
		moods := [...][3]string{
			{"good", "nice", "fantastic"},
			{"bad", "meh", "horrible"},
		}

		m := 0
		if args[1] == "negative" {
			m = 1
		}

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(moods[0]))

		fmt.Println(args[0] + " feels " + moods[m][n])
	}
}
