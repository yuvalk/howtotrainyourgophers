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
	for number := range numberChan {
		fmt.Println(number)
	}
	fmt.Println("Channel closed")
}

// END OMIT
