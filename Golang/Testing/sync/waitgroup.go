package main

import (
	"sync"
	"net/http"
)

func main(){
	wg := sync.WaitGroup{}
	urls := []string{
		"http://www.golang.org/",
        "http://www.google.com/",
	}

	for _, url := range urls{
		wg.Add(1)
		go func(){
			defer wg.Done()
			http.Get(url)
		}()
	}
	wg.Wait()
}