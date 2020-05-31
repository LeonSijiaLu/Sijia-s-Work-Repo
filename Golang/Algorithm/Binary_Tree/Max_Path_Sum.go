package main

import (
	"fmt"
)

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func getMax(a int, b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}

func getTripleMax(a int, b int, c int) int{
	return getMax(getMax(a, b), c)
}

func getPartialMaxPathSum(tree *BinaryTree) int{
	if tree == nil{return 0}
	return getMax(getPartialMaxPathSum(tree.Left), getPartialMaxPathSum(tree.Right)) + tree.Value
}

func MaxPathSum(tree *BinaryTree) int {
	return tree.Value + getPartialMaxPathSum(tree.Left) + getPartialMaxPathSum(tree.Right)
}

func main(){
	node4 := &BinaryTree{Value: 3, Left: nil, Right: nil}
	node3 := &BinaryTree{Value: 2, Left: nil, Right: nil}
	node2 := &BinaryTree{Value: -1, Left: node3, Right: node4}
	node1 := &BinaryTree{Value: -2, Left: node2, Right: nil}
	fmt.Println(MaxPathSum(node1))
}