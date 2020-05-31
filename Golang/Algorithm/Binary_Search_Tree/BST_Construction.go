package main

import "fmt"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(Value int) *BST {
	if Value >= tree.Value {
		if tree.Right != nil {
			tree.Right.Insert(Value)
		} else {
			tree.Right = &BST{Value: Value}
		}
	} else {
		if tree.Left != nil {
			tree.Left.Insert(Value)
		} else {
			tree.Left = &BST{Value: Value}
		}
	}
	return tree
}

func (tree *BST) Remove(Value int) *BST {
	if tree.Contains(Value) {
		tree.remove(Value, nil)
		return tree
	} else {
		return tree
	}
}

func (tree *BST) remove(Value int, parent *BST) {
	if Value < tree.Value{
		if tree.Left != nil{
			tree.Left.remove(Value, tree)	
		}
	}else if Value > tree.Value{
		if tree.Right != nil{
			tree.Right.remove(Value, tree)
		}
	}else{
		if tree.Left != nil && tree.Right != nil{
			successor := tree.Right.getSuccessor()
			tree.Value = successor.Value
			tree.Right.remove(successor.Value, tree)
		}else if parent == nil{ // Left is nil or Right is nil
			if tree.Right != nil{
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			}else if tree.Left != nil{
				tree.Value = tree.Left.Value
				tree.Left = tree.Left.Left
				tree.Right = tree.Left.Right
			}
		}else if parent.Left == tree{
			if tree.Left != nil{
				parent.Left = tree.Left
			}else{
				parent.Left = tree.Right
			}
		}else if parent.Right == tree{
			if tree.Right != nil{
				parent.Right = tree.Right
			}else{
				parent.Right = tree.Left
			}
		}
	}
}

func (tree *BST) getSuccessor() *BST {
	isDone := false
	for !isDone {
		isDone = true
		if tree.Left != nil {
			tree = tree.Left
			isDone = false
		}
	}
	return tree
}

func (tree *BST) Contains(Value int) bool {
	if tree.Value != Value {
		switch {
		case Value < tree.Value && tree.Left == nil: // leftmost
			return false
		case Value > tree.Value && tree.Right == nil: //Rightmost
			return false
		case Value < tree.Value && Value > tree.Left.Value:
			return false
		case Value > tree.Value && Value < tree.Right.Value:
			return false
		case Value < tree.Value && tree.Left != nil && Value < tree.Left.Value:
			return tree.Left.Contains(Value)
		case Value > tree.Value && tree.Right != nil && Value > tree.Right.Value:
			return tree.Right.Contains(Value)
		}
	}
	return true
}

func (tree *BST) InOrderTraversal(BST_array []int) {
	if tree.Left != nil {
		tree.Left.InOrderTraversal(BST_array)
	}
	fmt.Println(tree.Value)
	if tree.Right != nil {
		tree.Right.InOrderTraversal(BST_array)
	}
}

func NewBST(node_value int, values ...int) *BST {
	root := &BST{Value: node_value}
	for _, each_value := range values {
		root.Insert(each_value)
	}
	return root
}

func main() {
	root := NewBST(10, 5, 15, 5, 2, 14, 22)
	BST_array := []int{}
	root.InOrderTraversal(BST_array)
	//fmt.Println(1, root.Contains(1))
	//fmt.Println(5, root.Contains(5))
	//fmt.Println(10, root.Contains(10))
	//fmt.Println(15, root.Contains(15))
	//fmt.Println(20, root.Contains(20))
}
