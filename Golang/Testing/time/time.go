package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Duration(10))
	fmt.Println(time.Duration(10) * time.Second)
	fmt.Println(10 * time.Second)
}
