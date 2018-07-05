package main

import (
	"fmt"
	"time"
)

type pizza bool

// START OMIT
func main() {

	c := make(chan pizza)
	fmt.Println("Customer: Ordered pizza")

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Waiter: Pizza ready")
		c <- true
	}()

	myPizza := <-c
	if myPizza {
		fmt.Println("Customer: pizza received")
	}
}

// END OMIT
