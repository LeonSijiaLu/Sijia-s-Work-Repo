package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "Monitor 1")
	go watch(ctx, "Monitor 2")
	go watch(ctx, "Monitor 3")

	time.Sleep(5 * time.Second)
	fmt.Println("Time to cancel")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " is stopped")
			return
		default:
			fmt.Println(name, " is running")
			time.Sleep(time.Second)
		}
	}
}
