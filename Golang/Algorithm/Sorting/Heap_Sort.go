package main

import (
	"fmt"
)

func HeapSort(array []int) []int{
	for i := (len(array) - 1) / 2; i >= 0; i --{
		heapify(array, len(array), i)
	}
	var arr []int
	return sort(array, arr)
}

func sort(array []int, arr []int) []int{
	max := 0
	for len(array) != 0{
		array[0], array[len(array) - 1] = array[len(array) - 1], array[0]
		max = array[len(array) - 1]
		array = array[0:len(array) - 1]
		arr = append(arr, max)
		heapify(array, len(array), 0)
	}
	return arr
}

func heapify(array []int, length int, index int){
	if index >= length{
		return
	}
	c1 := index * 2 + 1
	c2 := index * 2 + 2
	max := index
	if c1 < length && array[c1] > array[max]{
		max = c1
	}
	if c2 < length && array[c2] > array[max]{
		max = c2
	}
	if max != index{
		array[max], array[index] = array[index], array[max]
		heapify(array, length, max)
	}
}

func main(){
	array := []int{2, 5, 3, 1, 10, 4}
	fmt.Println(HeapSort(array))
}