package main

import (
	"fmt"
)

const INT_MAX = 1<<32 -1 

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	var solutions []int
	num := make([]int, n + 1)
	for i := range num{
		num[i] = INT_MAX
	}
	num[0] = 0

	for i, _ := range denoms{
		for j, _ := range num{
			if denoms[i] < num[j]{ // if this coin is possible
				
			}
		}
	}
}

func main(){
	n := 6
	denoms := []int{1, 5}
	fmt.Println(NumberOfWaysToMakeChange(n, denoms))
}