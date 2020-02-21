package main

import "fmt"

func Insertion_Sort(array []int) []int {
	fmt.Println(array)
}

func main() {
	array := []int{123, 534, 213, 8768, 42, 213, 8560, 324, 756, 123, 987, 345, 867, 3423, 987, 978, 908, 32432, 1234, 9870}
	Insertion_Sort(array)
}
