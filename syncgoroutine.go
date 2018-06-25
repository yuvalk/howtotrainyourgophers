package main

import (
	"fmt"
	"time"
)

type pizza bool

// START OMIT
func main() {

	c := make(chan pizza)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Pizza ready")
		c <- true
	}()

	fmt.Println("Ordered pizza")
	myPizza := <-c
	if myPizza {
		fmt.Println("Pizza received")
	}
}

// END OMIT
