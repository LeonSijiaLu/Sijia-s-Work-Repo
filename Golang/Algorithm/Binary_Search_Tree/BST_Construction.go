package main

import "fmt"

type BST struct {
	value int
	left  *BST
	right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value >= tree.value {
		if tree.right != nil {
			tree.right.Insert(value)
		} else {
			tree.right = &BST{value: value}
		}
	} else {
		if tree.left != nil {
			tree.left.Insert(value)
		} else {
			tree.left = &BST{value: value}
		}
	}
	return tree
}

func (tree *BST) Remove(value int) *BST{
	
}

func (tree *BST) Contains(value int) bool {
	if tree.value != value {
		switch {
		case value < tree.value && tree.left == nil: // leftmost
			return false
		case value > tree.value && tree.right == nil: //rightmost
			return false
		case value < tree.value && value > tree.left.value:
			return false
		case value > tree.value && value < tree.right.value:
			return false
		case value < tree.value && tree.left != nil && value < tree.left.value:
			return tree.left.Contains(value)
		case value > tree.value && tree.right != nil && value > tree.right.value:
			return tree.right.Contains(value)
		}
	}
	return true
}

func (tree *BST) InOrderTraversal(BST_array []int) {
	if tree.left != nil {
		tree.left.InOrderTraversal(BST_array)
	}
	fmt.Println(tree.value)
	if tree.right != nil {
		tree.right.InOrderTraversal(BST_array)
	}
}

func NewBST(node_value int, values ...int) *BST {
	root := &BST{value: node_value}
	for _, each_value := range values {
		root.Insert(each_value)
	}
	return root
}

func main() {
	root := NewBST(10, 5, 15, 5, 2, 14, 22)
	BST_array := []int{}
	root.InOrderTraversal(BST_array)
	fmt.Println(1, root.Contains(1))
	fmt.Println(5, root.Contains(5))
	fmt.Println(10, root.Contains(10))
	fmt.Println(15, root.Contains(15))
	fmt.Println(20, root.Contains(20))
}
