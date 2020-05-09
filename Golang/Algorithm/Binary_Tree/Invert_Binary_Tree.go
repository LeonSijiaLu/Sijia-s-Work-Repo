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
	queue := []*BinaryTree{tree}
	for len(queue) > 0{
		curr := queue[0]
		queue = queue[1:]
		if curr == nil{
			continue
		}
		curr.Left, curr.Right = curr.Right, curr.Left
		queue = append(queue, curr.Left, curr.Right)
	}
}