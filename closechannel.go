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
	for {
		number, ok := <-pizzaChan
		if ok {
			fmt.Println(number)
		} else {
			fmt.Println("Oven closed")
			return
		}
	}
}

// END OMIT
