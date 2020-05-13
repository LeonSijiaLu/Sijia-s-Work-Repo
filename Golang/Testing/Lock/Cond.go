package main

import (
	"fmt"
	"sync"
	"time"
)

// Cond的作用并不是保证在同一时刻仅有一个线程访问某一个共享数据，
// 而是在对应的共享数据的状态发生变化时，通知其他因此而被阻塞的线程

func main(){
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	go func(){
		for {
			cond.L.Lock()
			for condition == 0{
				cond.Wait() // 等待信号
			}
			condition -- 
			fmt.Println("Consumer: ", condition)
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	for {
		time.Sleep(time.Second)
		cond.L.Lock()
		for condition == 3{
			cond.Wait()
		}
		condition ++
		fmt.Println("Producer: ", condition)
		cond.Signal()
		cond.L.Unlock()
	}
}