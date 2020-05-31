package main

import (
	"fmt"
)

// Do not edit the class below except for the buildHeap,
// siftDown, siftUp, peek, remove, and insert methods.
// Feel free to add new properties and methods to the class.
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	first := (len(array) - 2) / 2
	for current := first + 1; current >= 0; current --{
		h.siftDown(current, len(array) - 1)
	}
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	childOne := currentIndex * 2 + 1
	for childOne < endIndex{
		childTwo := -1
		if currentIndex * 2 + 2 != endIndex{
			childTwo = currentIndex * 2 + 2
		}
		indexSwap := childOne
		if childTwo > -1 && (*h)[childTwo] < (*h)[childOne]{
			indexSwap = childTwo
		}
		if (*h)[indexSwap] < (*h)[currentIndex]{
			(*h)[indexSwap], (*h)[currentIndex] = (*h)[currentIndex], (*h)[indexSwap]
			currentIndex = indexSwap
			childOne = currentIndex * 2 + 1
		}else{
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	cur := len((*h)) - 1
	parent := -1
	for cur > 0{
		if cur % 2 == 0{
			parent = (cur - 2) / 2
			indexSwap := -1
			if (*h)[cur] < (*h)[cur - 1]{
				indexSwap = cur
			}else{
				indexSwap = cur -1
			}
			if (*h)[indexSwap] < (*h)[parent]{
				(*h)[indexSwap], (*h)[parent] = (*h)[parent], (*h)[indexSwap]
				cur = parent
			}else{
				return
			}
		}else{
			parent = (cur - 1) / 2
			if (*h)[cur] < (*h)[parent]{
				(*h)[cur], (*h)[parent] = (*h)[parent], (*h)[cur]
				cur = parent
			}else{
				return
			}
		}
	}
}

func (h MinHeap) Peek() int {
	if len(h) == 0 {return -1}
	return h[0]
}

func (h *MinHeap) Remove() int {
	pop := (*h)[0]
	*h = (*h)[1:]
	if (*h)[0] > (*h)[1]{
		(*h)[0], (*h)[1] = (*h)[1], (*h)[0]
	}
	h.siftDown(0, len((*h)) - 1)
	return pop
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

func main(){
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	minheap := NewMinHeap(arr)
	minheap.Insert(76)
	minheap.Remove()
	minheap.Remove()
	minheap.Insert(87)
	fmt.Println(minheap)
}