package main

import (
	"fmt"
)

type MinHeap []int

func NewMinHeap(arr []int) *MinHeap{
	heap := MinHeap(arr)
	root := &heap
	root.BuildHeap(arr)
	return root
}

func (h *MinHeap) BuildHeap(arr []int){
	first := (len(arr) - 2) / 2
	for curr := first + 1; curr >= 0; curr --{
		h.SiftDown(curr, len(arr) - 1)
	}
}

func (h *MinHeap) Insert(item int){
	(*h) = append((*h), item)
	h.SiftUp()
}

func (h *MinHeap) SiftUp(){
	curr := len(*h) - 1
	for curr > 0{
		parent := (curr - 1) / 2
		if (*h)[curr] < (*h)[parent]{
			(*h)[curr], (*h)[parent] = (*h)[parent], (*h)[curr]
		}else{
			return
		}
		curr = (curr - 1) / 2
	}
}

func (h *MinHeap) SiftDown(curr int, length int){
	for curr < length{
		left_node := 2 * curr + 1
		if left_node > length {break}
		right_node := -1
		if left_node + 1 < length {right_node = left_node + 1}
		if right_node == -1{
			if (*h)[curr] > (*h)[left_node]{
				(*h)[curr], (*h)[left_node] = (*h)[left_node], (*h)[curr]
			}
		}else{
			if (*h)[left_node] < (*h)[right_node]{
				if (*h)[left_node] < (*h)[curr]{
					(*h)[curr], (*h)[left_node] = (*h)[left_node], (*h)[curr]
				}
			}else{
				if (*h)[right_node] < (*h)[curr]{
					(*h)[curr], (*h)[right_node] = (*h)[right_node], (*h)[curr]
				}
			}
		}
		curr = 2 * curr + 1
	}
}

func (h *MinHeap) Peek() int{
	return (*h)[0]
}

func main(){
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	heap := NewMinHeap(arr)
	fmt.Println(heap)

	heap.Insert(76)
	fmt.Println(heap)

	heap.Insert(-10)
	fmt.Println(heap)
}