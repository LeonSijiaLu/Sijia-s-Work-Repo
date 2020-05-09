package main

// How do we know if this array can form a cycle
// 在一次loop中
// 如果有一个点被命中两次，那么就有一个loop

func HasSingleCycle(array []int) bool {
	if len(array) == 1{
		return true
	}
	starting_point := 0
	curr_index := starting_point
	next_index := starting_point
	for i := 0; i < len(array); i ++{
		if curr_index == starting_point && i > 0{
			return false
		}
		jumps := array[curr_index]
		next_index = (curr_index + jumps) % len(array)
		if next_index < 0 {
			next_index = next_index + len(array)
		}
		curr_index = next_index
	}
	return curr_index == 0
}



