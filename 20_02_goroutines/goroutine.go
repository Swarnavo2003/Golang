package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when the function exits
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from Goroutine!")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// Start a new goroutine
	wg.Add(1) // Indicate that we are adding a new goroutine to wait for
	go sayHello(&wg)

	// Main function continues to execute
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from Main!")
		time.Sleep(150 * time.Millisecond)
	}

	// Wait for all goroutines to finish
	wg.Wait() // Blocks until the WaitGroup counter is zero
}