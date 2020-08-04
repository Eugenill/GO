package main

import (
	"context"
	"fmt"

	"time"
)

func main() {
	ctx := context.Background()
	c1 := make(chan string, 1)
	go func(ctx context.Context) {
		fmt.Println("<- ctx.Done()")
		time.Sleep(10 * time.Second)
		fmt.Println("result 1")
	}(ctx)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	printing(ctx)

	return
}
func printing(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		time.Sleep(2 * time.Second)
		fmt.Println("alive")
	}(ctx)
	cancel()
	return
}
