package main

import (
	"fmt"
)

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string{
	queue := []*Node{n}
	for len(queue) > 0{
		current := queue[0]
		queue = queue[1:]
		array = append(array, current.Name)
		for _, child := range current.Children{
			queue = append(queue, child)
		}
	}
	return array
}

func main(){
	K := &Node{Name: "K"}
	J := &Node{Name: "J"}
	I := &Node{Name: "I"}
	E := &Node{Name: "E"}
	F := &Node{Name: "F"}
	F.Children = append(F.Children, I)
	F.Children = append(F.Children, J)
	G := &Node{Name: "G"}
	G.Children = append(G.Children, K)
	H := &Node{Name: "H"}
	B := &Node{Name: "B"}
	B.Children = append(B.Children, E)
	B.Children = append(B.Children, F)
	C := &Node{Name: "C"}
	D := &Node{Name: "D"}
	D.Children = append(D.Children, G)
	D.Children = append(D.Children, H)
	A := &Node{Name: "A"}
	A.Children = append(A.Children, B)
	A.Children = append(A.Children, C)
	A.Children = append(A.Children, D)

	fmt.Println(A.BreadthFirstSearch([]string{}))
}