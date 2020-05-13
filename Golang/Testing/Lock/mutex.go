package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var mutex sync.Mutex
	mutex.Lock()
	for i := 1; i < 4; i++{
		go func(i int){
			fmt.Println("Locking ", i)
			mutex.Lock()
			fmt.Println("Mutex is locked ", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocking")
	mutex.Unlock()
	fmt.Println("Mutex is unlocked")
}