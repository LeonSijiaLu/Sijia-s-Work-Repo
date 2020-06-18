package main

import (
	"fmt"
)

type Node struct{
	name string
	neighbour string
	distance int
	goal_distance int
}

type queue_node struct{
	name string
	value int
}

func checkExistence(new_node string, visited []string) bool{
	for i := range visited{
		if visited[i] == new_node{
			return true
		}
	}
	return false
}

func removeDuplicates(queue []*queue_node) []*queue_node{
	counter := 0 
	for counter < len(queue) - 1{
		for i := counter + 1; i < len(queue); i ++{
			if queue[counter].name == queue[i].name{
				temp := queue[i + 1:]
				queue = queue[:i]
				queue = append(queue, temp...)
				break
			}
		}
		counter ++
	}
	return queue
}

func showQueue(queue []*queue_node){
	var arr []string
	for i := range queue{
		arr = append(arr, queue[i].name)
	}
	fmt.Println("Open Queue is :")
	fmt.Println(arr)
}

func findNodeValue(name string, edges[]*Node) int{
	for i := range edges{
		if name == edges[i].name{
			return edges[i].goal_distance
		}
	}
	return 0
}

func sort(queue []*queue_node) []*queue_node{
	for i := range queue{
		for j := range queue{
			if (queue[j].value > queue[i].value) || (queue[j].value == queue[i].value && queue[j].name > queue[i].name){
				queue[j], queue[i] = queue[i], queue[j]
			}
		}
	}
	return queue
}

func greedy_best_search(edges []*Node, starting string, ending string) ([]string, int){
	var queue []*queue_node
	queue = append(queue, &queue_node{name: starting, value: 78})

	var visited []string

	finished := false
	steps := 0
	for !finished{
		showQueue(queue)
		fmt.Println("Closed Queue is :")
		fmt.Println(visited)
		fmt.Println("=========================================")
		fmt.Println()
		pop := queue[0]
		queue = queue[1:]
		if !checkExistence(pop.name, visited){ // if this node not visited
			for i := range edges{
				if edges[i].name == pop.name && !checkExistence(edges[i].neighbour, visited){ // "S" == "S" and not visited
					queue = append(queue, &queue_node{name: edges[i].neighbour, value: findNodeValue(edges[i].neighbour, edges)})
				}
			}	
			visited = append(visited, pop.name)
		}
		queue = sort(queue)
		queue = removeDuplicates(queue)
		if pop.name == ending{
			finished = true
			steps = pop.value
		}
	}
	fmt.Println(visited, steps)
	return visited, steps
}

func main(){
	var edges []*Node
	/*edges = append(edges, &Node{name: "S", neighbour: "A", distance: 5, goal_distance: 5})
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
	edges = append(edges, &Node{name: "F", neighbour: "G", distance: 8, goal_distance: 6})*/

	edges = append(edges, &Node{name: "6", neighbour: "2", distance: 38, goal_distance: 60})
	edges = append(edges, &Node{name: "6", neighbour: "5", distance: 35, goal_distance: 60})
	edges = append(edges, &Node{name: "6", neighbour: "10", distance: 30, goal_distance: 60})
	edges = append(edges, &Node{name: "6", neighbour: "9", distance: 26, goal_distance: 60})
	edges = append(edges, &Node{name: "5", neighbour: "1", distance: 5, goal_distance: 75})
	edges = append(edges, &Node{name: "2", neighbour: "9", distance: 26, goal_distance: 32})
	edges = append(edges, &Node{name: "9", neighbour: "10", distance: 26, goal_distance: 35})
	edges = append(edges, &Node{name: "9", neighbour: "4", distance: 18, goal_distance: 35})
	edges = append(edges, &Node{name: "9", neighbour: "7", distance: 35, goal_distance: 35})
	edges = append(edges, &Node{name: "2", neighbour: "7", distance: 32, goal_distance: 32})
	edges = append(edges, &Node{name: "1", neighbour: "8", distance: 24, goal_distance: 78})
	edges = append(edges, &Node{name: "10", neighbour: "8", distance: 15, goal_distance: 57})
	edges = append(edges, &Node{name: "10", neighbour: "3", distance: 24, goal_distance: 57})
	edges = append(edges, &Node{name: "8", neighbour: "3", distance: 23, goal_distance: 60})
	edges = append(edges, &Node{name: "4", neighbour: "3", distance: 7, goal_distance: 30})

	edges = append(edges, &Node{name: "2", neighbour: "6", distance: 38, goal_distance: 32})
	edges = append(edges, &Node{name: "5", neighbour: "6", distance: 35, goal_distance: 75})
	edges = append(edges, &Node{name: "10", neighbour: "6", distance: 30, goal_distance: 57})
	edges = append(edges, &Node{name: "9", neighbour: "6", distance: 26, goal_distance: 35})
	edges = append(edges, &Node{name: "1", neighbour: "5", distance: 5, goal_distance: 78})
	edges = append(edges, &Node{name: "9", neighbour: "2", distance: 26, goal_distance: 35})
	edges = append(edges, &Node{name: "10", neighbour: "9", distance: 26, goal_distance: 57})
	edges = append(edges, &Node{name: "4", neighbour: "9", distance: 18, goal_distance: 30})
	edges = append(edges, &Node{name: "7", neighbour: "9", distance: 35, goal_distance: 0})
	edges = append(edges, &Node{name: "7", neighbour: "2", distance: 32, goal_distance: 0})
	edges = append(edges, &Node{name: "8", neighbour: "1", distance: 24, goal_distance: 60})
	edges = append(edges, &Node{name: "8", neighbour: "10", distance: 15, goal_distance: 60})
	edges = append(edges, &Node{name: "3", neighbour: "10", distance: 24, goal_distance: 37})
	edges = append(edges, &Node{name: "3", neighbour: "8", distance: 23, goal_distance: 37})
	edges = append(edges, &Node{name: "3", neighbour: "4", distance: 7, goal_distance: 37})

	greedy_best_search(edges, "1", "7")
}