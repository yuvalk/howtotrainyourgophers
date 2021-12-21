package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cpg1111/threadpool-go"
)

func main() {
	count := 0
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	tp, err := threadpool.New(ctx, cancel, 10000)
	if err != nil {
		fmt.Printf("threadpool err: %s\n", err)
	}
	tp.Start()

	for i := 0; i < 10000; i++ {
		tp.Exec(func() {
			count++
		})
	}

	fmt.Printf("end of threads creation count=%d\n", count)
}
