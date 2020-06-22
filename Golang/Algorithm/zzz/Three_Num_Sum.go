package main

import (
	"fmt"
	"sort"
)

func three_num_sums(arr []int, target int) [][]int{
	var res [][]int
	sort.Ints(arr)
	if arr[0] > target{
		return res
	}else{
		for i := 0; i < len(arr); i ++{
			left, right := i + 1, len(arr) - 1
			for left < right{
				curr_sum := arr[left] + arr[right] + arr[i]
				if curr_sum == target{
					res = append(res, []int{arr[i], arr[left], arr[right]})
					left ++
					right --
				}else if curr_sum < target{
					left ++
				}else if curr_sum > target{
					right --
				}
			}
		}
	}
	return res
}

func main(){
	arr := []int{12, 3, 1, 2, -6, 5, -8, 6}
	res := three_num_sums(arr, 0)
	fmt.Println(res)
}