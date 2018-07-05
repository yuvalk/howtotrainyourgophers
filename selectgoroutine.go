package main

import (
	"fmt"
	"time"
)

func customerGenerator() chan int {
	number := 0
	res := make(chan int)
	go func() {
		for {
			number++
			res <- number
		}
	}()
	return res
}

var pizzaGenerator = customerGenerator

// START OMIT
func main() {
	customerOneChan := customerGenerator()
	customerTwoChan := customerGenerator()
	pizzaReadyChan := pizzaGenerator()

	for {
		select {
		case <-pizzaReadyChan:
			fmt.Println("Pizza ready")
		case n := <-customerOneChan:
			fmt.Println("Order from customer one", n)
		case n := <-customerTwoChan:
			fmt.Println("Order from customer two", n)
		}
		time.Sleep(time.Second)
	}
}

// END OMIT
