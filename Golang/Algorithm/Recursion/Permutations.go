package main

import (
	"fmt"
	//"sort"
	//"strings"
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
	for index := range array{ // 1, 2, 3, 4
		newArray := make([]int, index)
		copy(newArray, array[:index])
		newArray = append(newArray, array[index + 1:]...) // get rid of the current number

		newPermutations := make([]int, len(currentPermutations))
		copy(newPermutations, currentPermutations)

		newPermutations = append(newPermutations, array[index])
		permutationsHelper(newArray, newPermutations, permutations)
	}
}

func main(){
	array := []int{1,2,3}
	fmt.Println(GetPermutations(array))
}