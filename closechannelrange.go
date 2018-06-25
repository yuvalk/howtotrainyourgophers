package main

import (
	"fmt"
	"time"
)

func numberGenerator() chan int {
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
	numberChan := numberGenerator()
	for {
		number, ok := <-numberChan
		if ok {
			fmt.Println(number)
		} else {
			fmt.Println("Channel closed")
			return
		}
	}
}

// END OMIT
