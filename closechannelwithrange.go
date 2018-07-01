package main

import (
	"fmt"
	"time"
)

func pizzaGenerator() chan int {
	res := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			res <- i
			time.Sleep(time.Second)
		}
		close(res)
	}()
	return res
}

// START OMIT

func main() {
	pizzaChan := pizzaGenerator()
	for pizza := range pizzaChan {
		fmt.Println("Pizza ", pizza)
	}
	fmt.Println("Channel closed")
}

// END OMIT
