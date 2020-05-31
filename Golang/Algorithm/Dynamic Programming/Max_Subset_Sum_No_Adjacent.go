package main

func getMax(a int, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}

// 如果取这个值：f(i) = f(i - 2) + f(i)
// 如果不取值：f(i) = f(i - 1)
// Max(取值或者不取值)

func MaxSubsetSumNoAdjacent(array []int) int{
	if len(array) == 0{return 0}
	if len(array) == 1{return array[0]}
	if len(array) == 2{return getMax(array[0], array[1])}
	arr := array
	arr[0] = array[0]
	arr[1] = getMax(array[0], array[1])
	for i := 2; i < len(array); i ++{
		arr[i] = getMax(array[i - 2] + array[i], array[i - 1])
	}
	return arr[len(arr) - 1]
}

func main(){
	array := []int{75, 105, 120, 75, 90, 135}
	MaxSubsetSumNoAdjacent(array)
}