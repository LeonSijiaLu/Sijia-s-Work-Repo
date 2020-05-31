package main

import (
	"fmt"
	"errors"
)

const (
	RED bool = true
	BLACK bool = false
	Left_Rotate bool = true
	Right_Rotate bool = false
)

type RBNode struct{
	value int
	color bool
	left *RBNode
	right *RBNode
	parent *RBNode
}

type RBTree struct{
	root *RBNode
}

func (rbnode *RBNode) rotate(isRotateLeft bool) (*RBNode, error){
	if rbnode == nil{
		return rbnode, nil
	}
	if !isRotateLeft && rbnode.left == nil{ // 右旋转，并且左节点不为空
		return rbnode, errors.New("Left cannot be nil in a right rotate")
	}else if isRotateLeft && rbnode.right == nil{
		return rbnode, errors.New("Right cannot be nil in a left rotate")
	}
	var root *RBNode
	var isleftnode bool
	parent := rbnode.parent
	if parent != nil{
		if parent.left == rbnode{
			isleftnode = true
		}else{
			isleftnode = false
		}
	}

	if isRotateLeft{
		grandson := rbnode.right.left
		rbnode.right.left = rbnode
		rbnode.parent = rbnode.right
		rbnode.right = grandson
	}else{
		grandson := rbnode.left.right
		rbnode.left.right = rbnode
		rbnode.parent = rbnode.left
		rbnode.left = grandson
	}

	if parent == nil{
		rbnode.parent.parent = nil
		root = rbnode.parent
	}else{
		if isleftnode{
			parent.left = rbnode.parent
		}else{
			parent.right = rbnode.parent
		}
		rbnode.parent.parent = parent
	}
	return root, nil
}
