package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sum_array := make([]int, 0)
	sum := 0
	fmt.Println(Depth_First_Search(root, sum_array, sum))
	return sum_array
}

func Depth_First_Search(root *BinaryTree, sum_array []int, sum int) []int {
	sum = sum + root.Value
	if root.Left != nil {
		sum_array = Depth_First_Search(root.Left, sum_array, sum)
	}
	if root.Right != nil {
		sum_array = Depth_First_Search(root.Right, sum_array, sum)
	}
	if root.Left == nil && root.Right == nil {
		sum_array = append(sum_array, sum)
	}
	return sum_array
}

func main() {
	node10 := &BinaryTree{Value: 10, Left: nil, Right: nil}
	node9 := &BinaryTree{Value: 9, Left: nil, Right: nil}
	node8 := &BinaryTree{Value: 8, Left: nil, Right: nil}
	node7 := &BinaryTree{Value: 7, Left: nil, Right: nil}
	node6 := &BinaryTree{Value: 6, Left: nil, Right: nil}
	node5 := &BinaryTree{Value: 5, Left: nil, Right: node10}
	node4 := &BinaryTree{Value: 4, Left: node8, Right: node9}
	node3 := &BinaryTree{Value: 3, Left: node6, Right: node7}
	node2 := &BinaryTree{Value: 2, Left: node4, Right: node5}
	node1 := &BinaryTree{Value: 1, Left: node2, Right: node3}
	BranchSums(node1)
}
