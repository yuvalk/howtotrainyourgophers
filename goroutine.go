package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	fmt.Println("Customer: ordered pizza")

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("Waiter: Pizza ready")
	}()
}

// END OMIT
