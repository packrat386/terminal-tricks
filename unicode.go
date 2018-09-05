package main

import (
	"fmt"
	"time"
)

// BEGIN CODE OMIT
var clock = []rune{'🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚', '🕛', '🕐', '🕑'}

func drawSpinners(i int) {
	fmt.Printf("\rworking: %c", clock[i%len(clock)])
}

// END CODE OMIT

func main() {
	status := make(chan int)
	go func() {
		for i := 0; i <= 50; i++ {
			status <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(status)
	}()

	for p := range status {
		drawSpinners(p)
	}

	// Start a new line for the next command prompt
	fmt.Printf("\r   done! %c\n", '😃')
}
