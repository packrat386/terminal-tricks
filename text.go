package main

import "fmt"

// BEGIN CODE OMIT
func reset() { fmt.Print("\033[0m") }

func main() {
	fmt.Printf("\033[1mguys")
	reset()

	fmt.Printf(" look at all the ")

	fmt.Printf("\033[3mcool")
	reset()

	fmt.Printf(" ways we can put ")

	fmt.Printf("\033[4memphasis")
	reset()

	fmt.Printf(" on text!\n")
}

// END CODE OMIT
