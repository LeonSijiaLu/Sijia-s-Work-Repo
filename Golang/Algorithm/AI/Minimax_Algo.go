package main

import (
	"fmt"
)

const INT_MAX = 1<<32 -1 
var arr []string

type Node struct{
	children []*Node
	uppper int
	bottom int
	node_type bool // 0 = min, 1 = max
	value int
	symbol string
}

func MinMax_Search(root *Node){
	if len(root.children) != 0{
		for i := range root.children{
			arr = append(arr, root.children[i].symbol)
			MinMax_Search(root.children[i])
		}
	}
}

func main(){	
	A := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 0, symbol: "A"}

	B := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: false, value: 0, symbol: "B"}
	C := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: false, value: 0, symbol: "C"}

	D := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 3, symbol: "D"}
	E := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 6, symbol: "E"}
	F := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 0, symbol: "F"}
	G := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 0, symbol: "G"}
	H := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 4, symbol: "H"}

	I := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: false, value: 0, symbol: "I"}
	J := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: false, value: 5, symbol: "J"}
	K := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: false, value: 6, symbol: "K"}
	L := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: false, value: 10, symbol: "L"}

	M := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 1, symbol: "M"}
	N := &Node{children: nil, uppper: INT_MAX, bottom: -INT_MAX, node_type: true, value: 4, symbol: "N"}

	A.children = append(A.children, B)
	A.children = append(A.children, C)

	B.children = append(B.children, D)
	B.children = append(B.children, E)

	C.children = append(C.children, F)
	C.children = append(C.children, G)
	C.children = append(C.children, H)

	F.children = append(F.children, I)
	F.children = append(F.children, J)

	G.children = append(G.children, K)
	G.children = append(G.children, L)

	I.children = append(I.children, M)
	I.children = append(I.children, N)

	MinMax_Search(A)

	fmt.Println(arr)
}