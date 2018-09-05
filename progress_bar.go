package main

import (
	"fmt"
	"strings"
	"time"
)

// BEGIN CODE OMIT
func drawBar(percent int) {
	cols := 20
	prog := cols * percent / 100.0
	bar := strings.Repeat("=", prog)
	if prog != 20 {
		bar = bar + ">"
	}

	fmt.Printf("\rdemo progress: %3[1]d%% |%-[3]*[2]s|",
		percent, bar, cols)
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
		drawBar(2 * p)
	}

	// Start a new line for the next command prompt
	fmt.Println()
}
