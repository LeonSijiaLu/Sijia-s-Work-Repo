package main

import (
	"fmt"
)

// 如果不拿, f(i) = f(i - 1)
// 如果拿 f(i) = f(i - 1) + f(i)

func getMax(a int, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}

func KnapsackProblem(items [][]int, capacity int) []interface{} {
	
}

func main(){
	capacity := 10
	items := [][]int{
		{1, 2}, // value, weight
		{4, 3},
		{5, 6},
		{6, 7},
	}
	fmt.Println(KnapsackProblem(items, capacity))
}