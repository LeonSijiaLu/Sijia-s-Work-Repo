package main

import (
	"fmt"
	"math"
)

func BinarySearch(array []int, target int) int { // return the index of the target value in the array
	mid_index := int(math.Floor(float64(len(array)-1) / float64(2)))
	if len(array) == 0 {
		return -1
	}
	if array[mid_index] < target {
		array_right_half := array[mid_index+1:]
		return BinarySearch(array_right_half, target)
	} else if array[mid_index] > target {
		array_left_half := array[0:mid_index]
		return BinarySearch(array_left_half, target)
	} else {
		return array[mid_index]
	}
}

func BinarySearch_Index(array []int, target int, left int, right int) int {
	mid_point := int(math.Floor(float64(left+right) / float64(2)))
	if len(array) == 0 {
		return -1
	}
	if left > right || left > len(array)-1 || right > len(array)-1 {
		return -1
	}
	if array[mid_point] < target {
		return BinarySearch_Index(array, target, mid_point+1, right)
	} else if array[mid_point] > target {
		return BinarySearch_Index(array, target, left, mid_point-1)
	} else {
		return mid_point
	}
}

func main() {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	fmt.Println(BinarySearch_Index(array, 79, 0, len(array)-1))
}
