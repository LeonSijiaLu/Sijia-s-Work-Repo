package main

import (
	"fmt"
)

func Selection_Sort(array []int) []int {
	counter := 0
	for counter != len(array)-1 {
		min_v := array[counter]
		min_index := counter
		for i := counter; i <= len(array)-1; i++ {
			if array[i] < min_v {
				min_v = array[i]
				min_index = i
			}
		}
		array[counter], array[min_index] = array[min_index], array[counter]
		counter++
	}
	return array
}

func main() {
	array := []int{123, 534, 213, 8768, 42, 25, 213, 8560, 324, 756, 123, 987, 345, 867, 3423, 987, 978, 908, 32432, 1234, 9870}
	fmt.Println(Selection_Sort(array))
	// Put the smallest number on the left
	// Put the second smallest number on the second left
}
