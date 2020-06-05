
package main

import (
	"fmt"
)

type Node struct{
	name string
	neighbour string
	distance int
}

type queue_node struct{
	name string
	value int
}

func greedy_best_search(){
	
}

func main(){
	var edges []*Node
	edges = append(edges, &Node{name: "S", neighbour: "A", distance: 5, goal_distance: 5})
	edges = append(edges, &Node{name: "S", neighbour: "B", distance: 9, goal_distance: 5})
	edges = append(edges, &Node{name: "S", neighbour: "D", distance: 6, goal_distance: 5})
	edges = append(edges, &Node{name: "A", neighbour: "G", distance: 9, goal_distance: 7})
	edges = append(edges, &Node{name: "A", neighbour: "B", distance: 3, goal_distance: 7})
	edges = append(edges, &Node{name: "B", neighbour: "C", distance: 1, goal_distance: 3})
	edges = append(edges, &Node{name: "B", neighbour: "A", distance: 2, goal_distance: 3})
	edges = append(edges, &Node{name: "C", neighbour: "S", distance: 6, goal_distance: 4})
	edges = append(edges, &Node{name: "C", neighbour: "G", distance: 5, goal_distance: 4})
	edges = append(edges, &Node{name: "C", neighbour: "F", distance: 7, goal_distance: 4})
	edges = append(edges, &Node{name: "D", neighbour: "C", distance: 2, goal_distance: 6})
	edges = append(edges, &Node{name: "D", neighbour: "E", distance: 2, goal_distance: 6})
	edges = append(edges, &Node{name: "D", neighbour: "S", distance: 1, goal_distance: 6})
	edges = append(edges, &Node{name: "E", neighbour: "G", distance: 7, goal_distance: 5})
	edges = append(edges, &Node{name: "F", neighbour: "D", distance: 2, goal_distance: 6})
	edges = append(edges, &Node{name: "F", neighbour: "G", distance: 8, goal_distance: 6})

	greedy_best_search(edges)
}