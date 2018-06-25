package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("Pizza ready")
	}()
	fmt.Println("Pizza ordered")

}

// END OMIT
