package main

import (
	"fmt"
)

// if f[i-1] > 0
// then f[i] = f[i-1] + nums[i]
// else f[i-1]

func getMax(a int, b int) int{
	if a > b {
		return a
	}else{
		return b
	}
}

func KadanesAlgorithm(array []int) int {
	if len(array) == 1{return array[0]}
	max := array[0]
	for i := 1; i < len(array); i ++{
		array[i] = getMax(array[i], array[i] + array[i-1])
		if array[i] > max{
			max = array[i]
		}
	}
	return max
}

func main(){
	arr := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	fmt.Println(KadanesAlgorithm(arr))
}