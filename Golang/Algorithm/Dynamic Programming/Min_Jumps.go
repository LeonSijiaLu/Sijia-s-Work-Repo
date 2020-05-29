package main

import (
	"fmt"
)

// 如果走了
// 我们肯定选剩下范围里最大的
//
// 如果没走
// 直接下一个

func getMaxNumIndex(arr []int) int{
	max := 0
	index := 0
	for i := range arr{
		if arr[i] >= max{
			max = arr[i]
			index = i
		}
	}
	end := len(arr) - 1
	if arr[end] >= arr[index] - (end - index){
		return end
	}
	return index
}

func MinNumberOfJumps(array []int) int {
	if len(array) == 1{return 0}
	steps := 0
	for i := 0; i < len(array); i ++{
		steps ++
		if i + array[i] < len(array) - 1{
			i = i + getMaxNumIndex(array[i + 1: i + array[i] + 1])
		}else{
			return steps
		}
	}
	return steps
}

func main(){
	array := []int{3, 12, 2, 1, 2, 3, 7, 1, 1, 1, 3, 2, 3, 2, 1, 1, 1, 1}
	fmt.Println(MinNumberOfJumps(array))
}