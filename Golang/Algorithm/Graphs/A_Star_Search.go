// 思想：
// A * Search
// 首先 A * Search 是 Informed Search, 我们不仅知道每条边的长度，并且可以根据长度选择最优路线。最重要的是我们知道每个node之间的直接距离或者说相关联程度
// 1. 和 Uniform Cost Search 一样， A * Search 使用了 Priority Queue
// 2. 我们以BFS的方式进行遍历
// 比如我们有 (A, 5), (D, 6)
// 我们会先 pop (A, 5), 之后把 (A, 5) 周边的 (B, 8), (H, 14) push 进入 Priority Queue
// 这样 Priority Queue 内部就有：(D, 6), (B, 8), (H, 14)
// 之后我们 pop (D, 6), 然后重复
// **********     **********     **********     **********
// 3. 但是 A * Search 有一点很特殊，那就是我们在 push (B, x) 进入 Priority Queue 的时候，如何计算 x
// x = B 自己的 goal_distance + B 到 starting point 的 距离
//         S 0
//       / 2
//     A 3      那么A的值 = 2 + 3 = 5
//    / 3
//   B 4        那么B的值 = 2 + 3 + 4 = 9

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
	starting_distance int
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
			if (queue[j].value > queue[i].value) || (queue[j].value == queue[i].value && queue[j].name > queue[i].name){
				queue[j], queue[i] = queue[i], queue[j]
			}
		}
	}
	return queue
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

func A_Star_Search(edges []*Node, starting string, ending string) ([]string, int){
	var queue []*queue_node
	queue = append(queue, &queue_node{name: starting, value: 5, starting_distance: 0})

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
					queue = append(queue, &queue_node{name: edges[i].neighbour, value: edges[i].distance + pop.starting_distance + findNodeValue(edges[i].neighbour, edges), starting_distance: edges[i].distance + pop.starting_distance})
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

	A_Star_Search(edges, "1", "7")
}