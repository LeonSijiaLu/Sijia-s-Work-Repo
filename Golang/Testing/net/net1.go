package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println("False url")
	}
	fmt.Println(conn)
}
