package main

import (
	"fmt"
	"time"
)

// START OMIT

func numberGenerator(number int) chan int {
	res := make(chan int)
	go func() {
		res <- number
	}()
	return res
}

func main() {
	fmt.Println("Before wait")
	numChan := numberGenerator(0)
	time.Sleep(time.Second)
	value := <-numChan
	fmt.Println("Received ", value)
}

// END OMIT
