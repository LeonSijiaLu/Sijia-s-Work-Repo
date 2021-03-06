package main

import (
	"fmt"
)

func GetPermutations(array []int) [][]int {
	permutations := [][]int{}
	permutationsHelper(array, []int{}, &permutations)
	return permutations
}

func permutationsHelper(array []int, currentPermutations []int, permutations *[][]int){
	if len(array) == 0 && len(currentPermutations) != 0{
		*permutations = append(*permutations, currentPermutations)
		return
	}
	for i := range array{
		newArr := make([]int, len(array))
		copy(newArr, array[:i])
		newArr = append(newArr, array[i + 1:]...)
		newPermutations := make([]int, len(currentPermutations))
		copy(newPermutations, currentPermutations)
		newPermutations = append(newPermutations, array[i])
		permutationsHelper(newArr, newPermutations, permutations)
	}
}

func main(){
	arr := []int{1, 2, 3}
	fmt.Println(GetPermutations(arr))
}