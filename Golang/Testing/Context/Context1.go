package main

import (
	"context" //可以跟踪go routine的方案
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// context.Background返回一个空的context，这个空的Context位于整个Context的根节点
	// context.WithCancel()创建一个可以取消的子context
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Finished it !")
				return
			default:
				fmt.Println("Still running, dude")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("Time to finish this")
	cancel()
	time.Sleep(2 * time.Second)
}
