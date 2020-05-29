package main

import (
	"fmt"
)

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return getDepth(root, 0)
}

func getDepth(root *BinaryTree, height int) int{
	left_length := 0
	right_length := 0
	if root.Left != nil{
		left_length = getDepth(root.Left, height + 1)	
	}
	if root.Right != nil{
		right_length = getDepth(root.Right, height + 1)
	}
	if root.Left == nil && root.Right == nil{
		return height + left_length + right_length
	}
	return height + left_length + right_length
}

func main(){
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
	fmt.Println(NodeDepths(node1))
}
