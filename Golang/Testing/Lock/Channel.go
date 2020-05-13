package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int, 3)
	v := 0

	go func(){
		for {
			fmt.Println("Consumer: ", <-ch)
		}
	}()

	for{
		v++
		ch <- v
		fmt.Println("Producer: ", v)
		time.Sleep(time.Second)
	}
}