package main

import (
	"fmt"
	"time"
)

func customerGenerator(number int) chan int {
	res := make(chan int)
	go func() {
		for {
			res <- number
		}
	}()
	return res
}

// START OMIT
func main() {
	customerOneChan := customerGenerator(0)
	customerTwoChan := customerGenerator(1)
	pizzaReadyChan := customerGenerator(2)

	for {
		select {
		case <-pizzaReadyChan:
			fmt.Println("Pizza ready")
		case <-customerOneChan:
			fmt.Println("Customer one")
		case <-customerTwoChan:
			fmt.Println("Customer two")
		}
		time.Sleep(time.Second)
	}
}

// END OMIT
