package main

import (
	"fmt"
	"time"
)

type pizza bool

// START OMIT
func main() {
	pizzaReady := make(chan pizza, 1) // Note the buffered channel
	go func() {
		time.Sleep(10 * time.Second)
		pizzaReady <- true
	}()

	timeout := time.After(3 * time.Second)
	fmt.Println("Waiting for pizza")
	select {
	case <-pizzaReady:
		fmt.Println("Pizza ready")
	case <-timeout:
		fmt.Println("Too late, I am leaving")
	}
}

// END OMIT
