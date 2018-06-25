package main

import (
	"fmt"
	"time"
)

type order int

func makePizza(orderChan chan order) {
	for o := range orderChan {
		time.Sleep(time.Second)
		fmt.Printf("Pizza %d ready\n", o)
	}
}

func customer(id int) chan order {
	res := make(chan order)
	go func() {
		fmt.Printf("ordered %d\n", id)
		res <- order(id)
	}()
	return res
}

// START OMIT
func main() {
	pizzaMakerChan := make(chan order)
	c1 := customer(1)
	c2 := customer(2)
	c3 := customer(3)

	go makePizza(pizzaMakerChan)

	fmt.Println("Waiting for orders")
	for {
		select {
		case o := <-c1:
			pizzaMakerChan <- o
		case o := <-c2:
			pizzaMakerChan <- o
		case o := <-c3:
			pizzaMakerChan <- o
		case <-time.After(5 * time.Second):
			break
		}
	}
}

// END OMIT
