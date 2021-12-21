package main

import (
	"fmt"
)

// START OMIT
func main() {

	fmt.Println("Customer: ordered pizza")

	for i := 0; i < 50000; i++ {
		go func() {
		}()
	}
}

// END OMIT
