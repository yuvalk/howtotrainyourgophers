package main

import (
	"fmt"
	"time"
)

// START OMIT

func sum(inputChan <-chan int, outputChan chan<- int) {
	arg := <-inputChan
	outputChan <- arg * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)
	go sum(in, out)

	in <- 12
	res := <-out
	fmt.Println("Result ", res)

}

// END OMIT
