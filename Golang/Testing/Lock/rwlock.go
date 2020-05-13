package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	fmt.Println("WaitGroup ", wg)

	for i := 0; i < 10; i ++{
		wg.Add(1)
		go func(i int){
			fmt.Println(i, wg)
			time.Sleep(2 * time.Second)
			wg.Done()
			fmt.Println(i, wg)
		}(i)
	}

	wg.Wait()
	fmt.Println("Finished")
}

