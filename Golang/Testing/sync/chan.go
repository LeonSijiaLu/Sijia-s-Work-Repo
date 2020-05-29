package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	go func(){
		fmt.Println("Starting")
		time.Sleep(time.Second)
		ch <- 1
	}()
	<- ch
}