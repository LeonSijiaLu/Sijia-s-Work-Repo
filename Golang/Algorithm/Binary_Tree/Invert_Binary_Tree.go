package main

import (
	"fmt"
)

type BinaryTree struct{
	Value int
	Left *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree(){
	tree.Left, tree.Right = tree.Right, tree.Left
	if tree.Left != nil{
		tree.Left.InvertBinaryTree()
	}
	if tree.Right != nil{
		tree.Right.InvertBinaryTree()
	}
}

func (tree *BinaryTree) DFS(){
	fmt.Println(tree.Value)
	if tree.Left != nil{
		tree.Left.DFS()
	}
	if tree.Right != nil{
		tree.Right.DFS()
	}
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
	node1.InvertBinaryTree()
	node1.DFS()
}
