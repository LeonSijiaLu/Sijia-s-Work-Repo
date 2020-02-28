package main

import "fmt"

func ThreeNumberSum(array []int, target int) [][]int {
	result_array := [][]int{}
	sorted_array := InsertionSort(array)
	for _, num1 := range sorted_array {
		for _, num2 := range sorted_array {
			if num2 != num1 && num2 > num1 {
				for _, num3 := range sorted_array {
					if num3 != num2 && num3 != num1 && num3 > num2 && num3 > num1 {
						if num1+num2+num3 == target {
							result_array = append(result_array, []int{num1, num2, num3})
						}
					}
				}
			}
		}
	}
	return result_array
}

func InsertionSort(array []int) []int {
	counter := 1
	for counter != len(array) {
		for i := counter; i > 0; i-- {
			if array[i] < array[i-1] {
				array[i], array[i-1] = array[i-1], array[i]
			}
		}
		counter++
	}
	return array
}

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	fmt.Println(ThreeNumberSum(array, 0))
}
