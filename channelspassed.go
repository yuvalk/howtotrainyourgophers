package main

import (
	"fmt"
	"time"
)

// START OMIT

type order struct {
	pizza      string
	pizzaReady chan bool
}

func makePizza(orderChan chan order) {
	fmt.Println("Waiting for orders")
	for o := range orderChan {
		time.Sleep(time.Second)
		fmt.Printf("Pizza %s ready\n", o.pizza)
		o.pizzaReady <- true
	}
}

func customer(pizza string, orderChan chan order) {
	fmt.Printf("ordered %s\n", pizza)
	pizzaReadyChan := make(chan bool)
	orderChan <- order{pizza, pizzaReadyChan}
	<-pizzaReadyChan
	fmt.Printf("Eating %s\n", pizza)
}

// main links 3 customers to one pizza maker ...

// END OMIT

func main() {
	pizzaMakerChan := make(chan order)

	go makePizza(pizzaMakerChan)

	go customer("margherita", pizzaMakerChan)
	go customer("bomba", pizzaMakerChan)
	go customer("quattro formaggi", pizzaMakerChan)

	time.Sleep(10 * time.Second)
}
