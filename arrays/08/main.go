package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
		//{"Albert", "Isaac", "Stephen", "Marie", "Charles"},
		//{"Einstein", "Newton", "Hawking", "Curie", "Darwin"},
		//{"time", "apple", "blackhole", "radium", "fittest"},
	}
	//firstNames := [...]string{"Albert", "Isaac", "Stephen", "Marie", "Charles"}
	//surNames := [...]string{"Einstein", "Newton", "Hawking", "Curie", "Darwin"}
	//nickNames := [...]string{"time", "apple", "blackhole", "radium", "fittest"}

	fmt.Printf("%-10v %-10v %-10v\n", "First Name", "Last Name", "Nickname")
	fmt.Printf("%s\n", strings.Repeat("=", 30))

	for i := range names {
		n := names[i]
		fmt.Printf("%-10v %-10v %-10v\n", n[0], n[1], n[2])
	}
}
