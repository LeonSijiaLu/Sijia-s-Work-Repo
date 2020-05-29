package main

import (
	"fmt"
)

const INT_MAX = 1<<32 -1 

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	num := make([]int, n + 1)
	for i := range num{
		num[i] = INT_MAX
	}
	num[0] = 0
	for _, denom := range denoms{ // 1 yuan, 3 yuan, 5 yuan
		for j := range num{  // smallest coins
			if denom <= j {
				num[j] = min(num[j], num[j - denom] + 1)
			}
		}
	}
	if num[n] != INT_MAX{
		return num[n]
	}
	return -1
}

func min(arg int, data ...int) int{
	cur := arg
	for _, num := range data{
		if num < cur{
			cur = num
		}
	}
	return cur
}

func main(){
	n := 9
	denoms := []int{3, 5}
	fmt.Println(MinNumberOfCoinsForChange(n, denoms))
}