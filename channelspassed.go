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

func pizzaMaker(orderChan chan order) {
	fmt.Println("Pizzamaker : Waiting for orders")
	for o := range orderChan {
		time.Sleep(time.Second)
		fmt.Printf("Pizzamaker : Pizza %s ready\n", o.pizza)
		o.pizzaReady <- true
	}
}

func customer(pizza string, orderChan chan order) {
	fmt.Printf("Customer : ordered %s\n", pizza)
	pizzaReadyChan := make(chan bool)
	orderChan <- order{pizza, pizzaReadyChan}
	<-pizzaReadyChan
	fmt.Printf("Customer : Eating %s\n", pizza)
}

// main links 3 customers to one pizza maker ...

// END OMIT

func main() {
	pizzaMakerChan := make(chan order)

	go pizzaMaker(pizzaMakerChan)

	go customer("margherita", pizzaMakerChan)
	go customer("bomba", pizzaMakerChan)
	go customer("quattro formaggi", pizzaMakerChan)

	time.Sleep(10 * time.Second)
}
