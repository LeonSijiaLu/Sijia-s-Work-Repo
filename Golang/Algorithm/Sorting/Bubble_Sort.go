package main

import "fmt"

func Bubble_Sort_Ascend(array []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}
		counter++
	}
	fmt.Println(array)
	return array
}

func Bubble_Sort_Desc(array []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] < array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}
		counter++
	}
	fmt.Println(array)
	return array
}

func main() {
	array := []int{123, 534, 213, 8768, 42, 213, 8560, 324, 756, 123, 987}
	Bubble_Sort_Ascend(array)
	Bubble_Sort_Desc(array)
	// Put the largest at the back
	// Put the second largest at the second back
}
