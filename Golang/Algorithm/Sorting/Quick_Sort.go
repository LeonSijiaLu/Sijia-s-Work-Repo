package main

// 1. Choose a pivot
// 2. If itemFromLeft (first larger than pivot item) and itemFromRight (first smaller than pivot item)
// 3. Swap itemFromLeft with itemFromRight
// 4. If itemFromLeft's index is bigger than itemFromRight's index, then done this loop
// 5. Then run the rest

func Partition(array []int, low_bound int, high_bound int) int{
	pivot := array[0]
	left := low_bound + 1
	right := high_bound
	left_right_meet := false
	for{
		if left_right_meet == true{
			break;
		}
		if array[left] <= pivot { // increment left if correct
			left = left + 1
		}else if array[left] > pivot && array[right] < pivot{
			array[left], array[right] = array[right], array[left]
		}else if array[right] >= pivot{ // decrement right if correct
			right = right - 1
		}
		if right < left {
			left_right_meet = true
			array[0], array[right] = array[right], array[0]
		}
	}
	return right
}

func QuickSort(array []int) []int{
	if len(array) <= 1{
		return array
	}
	pivot := Partition(array, 0, len(array)-1)
	QuickSort(array[0: pivot])
	QuickSort(array[pivot + 1: len(array)])
	return array
}