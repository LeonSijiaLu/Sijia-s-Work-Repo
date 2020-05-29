package main

import (
	"time"
	"fmt"
	"context"
)

func PrintTask(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("A man")
		}
	}
}

func main(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go PrintTask(ctx)
	go PrintTask(ctx)
	go PrintTask(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("Exit")
}