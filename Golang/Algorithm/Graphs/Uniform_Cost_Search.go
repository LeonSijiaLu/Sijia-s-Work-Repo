// 思想：
// Uniform Cost Search
// 首先 Uniform Cost Search 是 Informed Search, 也就是说我们知道每条边的长度，并且可以根据长度选择最优路线
// 1. Uniform Cost Search 使用了 Priority Queue
// 2. 我们以BFS的方式进行遍历
// 比如我们有 (A, 5), (D, 6)
// 我们会先 pop (A, 5), 之后把 (A, 5) 周边的 (B, 8), (H, 14) push 进入 Priority Queue
// 这样 Priority Queue 内部就有：(D, 6), (B, 8), (H, 14)
// 之后我们 pop (D, 6), 然后重复
// 3. 整个Uniform Cost Search 直到我们 pop 出理想的stage才算结束

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

func checkExistence(new_node string, visited []string) bool{
	for i := range visited{
		if visited[i] == new_node{
			return true
		}
	}
	return false
}

func sort(queue []*queue_node) []*queue_node{
	for i := range queue{
		for j := range queue{
			if queue[j].value > queue[i].value{
				queue[j], queue[i] = queue[i], queue[j]
			}
		}
	}
	return queue
}

func uniform_cost_search(edges []*Node) ([]string, int){
	var queue []*queue_node
	queue = append(queue, &queue_node{name: "S", value: 0})

	var visited []string

	finished := false
	steps := 0
	for !finished{
		pop := queue[0]
		queue = queue[1:]
		if !checkExistence(pop.name, visited){ // if this node not visited
			for i := range edges{
				if edges[i].name == pop.name && !checkExistence(edges[i].neighbour, visited){ // "S" == "S" and not visited
					queue = append(queue, &queue_node{name: edges[i].neighbour, value: edges[i].distance + pop.value})
				}
			}	
			visited = append(visited, pop.name)
		}
		queue = sort(queue)
		if pop.name == "G"{
			finished = true
			steps = pop.value
		}
	}
	return visited, steps
}

func main(){
	var edges []*Node
	edges = append(edges, &Node{name: "S", neighbour: "A", distance: 5})
	edges = append(edges, &Node{name: "S", neighbour: "B", distance: 9})
	edges = append(edges, &Node{name: "S", neighbour: "D", distance: 6})
	edges = append(edges, &Node{name: "A", neighbour: "G", distance: 9})
	edges = append(edges, &Node{name: "A", neighbour: "B", distance: 3})
	edges = append(edges, &Node{name: "B", neighbour: "C", distance: 1})
	edges = append(edges, &Node{name: "B", neighbour: "A", distance: 2})
	edges = append(edges, &Node{name: "C", neighbour: "S", distance: 6})
	edges = append(edges, &Node{name: "C", neighbour: "G", distance: 5})
	edges = append(edges, &Node{name: "C", neighbour: "F", distance: 7})
	edges = append(edges, &Node{name: "D", neighbour: "C", distance: 2})
	edges = append(edges, &Node{name: "D", neighbour: "E", distance: 2})
	edges = append(edges, &Node{name: "D", neighbour: "S", distance: 1})
	edges = append(edges, &Node{name: "E", neighbour: "G", distance: 7})
	edges = append(edges, &Node{name: "F", neighbour: "D", distance: 2})
	edges = append(edges, &Node{name: "F", neighbour: "G", distance: 8})


	fmt.Println(uniform_cost_search(edges))
}