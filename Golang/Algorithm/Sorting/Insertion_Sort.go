package main

func InsertionSort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	counter := 1
	for {
		for i := counter; i > 0; i-- {
			if array[i] < array[i-1] {
				array[i-1], array[i] = array[i], array[i-1]
			}
		}
		counter++
		if counter == len(array) {
			break
		}
	}
	return array
}

//func main() {
//	array := []int{123, 534, 213, 8768, 42, 213, 8560, 324, 756, 123, 987}
//	fmt.Println(InsertionSort(array))
	// Round 1
	// 534 compares with 123, good
	//
	// Round 2
	// 213 compares with 534, swap
	// 213 compares with 123, good
	//
//}
