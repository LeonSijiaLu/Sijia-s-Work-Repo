package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (node *Node) newNode(value int) {
	node.value = value
	node.left = nil
	node.right = nil
}

func (node *Node) Insert(value int) error {
	if node == nil {
		return errors.New("Please construct a tree first")
	}

	switch {
	case value == node.value:
		return nil
	case value < node.value:
		if node.left == nil {
			node.left = &Node{value: value}
			return nil
		}
		node.left.Insert(value)
	case value > node.value:
		if node.right == nil {
			node.right = &Node{value: value}
		}
		node.right.Insert(value)
	}
	return nil
}

func (node *Node) find(value int) (int, bool) {
	switch {
	case value == node.value:
		fmt.Println(value, " is in the tree")
		return value, true
	case value < node.value:
		node.left.find(value)
	case value > node.value:
		node.right.find(value)
	}
	return -2, false
}

func (node *Node) InOrderTraverse() {
	if node == nil {
		return
	}
	node.left.InOrderTraverse()
	fmt.Println(node.value)
	node.right.InOrderTraverse()
}

func main() {
	root := new(Node)
	root.newNode(15)
	root.Insert(30)
	root.Insert(10)
	root.Insert(20)
	root.Insert(40)
	root.Insert(50)
	root.Insert(60)
	root.Insert(70)
	root.Insert(1)
	root.Insert(5)
	root.Insert(10)
	root.InOrderTraverse()
	root.find(10)
}
