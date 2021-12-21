package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	cow := 4
	processes := 0
	for i := 0; i < 10000; i++ {
		id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if id == 0 {
			cow++
			fmt.Println("In child:", id, cow, processes)
			os.Exit(0)
		} else {
			processes++
			fmt.Println("In parent:", id, cow, processes)
		}
	}
}
