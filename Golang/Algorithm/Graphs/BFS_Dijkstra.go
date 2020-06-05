package main

import (
	"fmt"
)

const MAX_INT = int(^uint(0) >> 1)

func sort(queue [][]int) [][]int{
	for i := range queue{
		for j := range queue{
			if queue[j][1] > queue[i][1]{
				queue[j], queue[i] = queue[i], queue[j]
			}
		}
	}
	return queue
}

func checkValid(res [][]int, check []int) [][]int{
	if len(res) == 0{
		res = append(res, check)
		return res
	}else{
		flag := false
		for i := range res{
			if res[i][0] == check[0]{ // same node
				flag = true
				if res[i][1] > check[1]{
					res[i][1] = check[1]
					return res
				}
			}
		}
		if flag == false{
			res = append(res, check)
		}
		return res
	}
	return res
}

func removeDuplicates(queue [][]int, pop []int) [][]int{
	for i := range queue{
		if pop[0] == queue[i][0]{ // same node
			queue[i][1] = MAX_INT
		}
	}
	return queue
}

func bfs_dijkstra(n int, edges [][]int, graph [][][]int) [][]int{
	var res [][]int
	var queue [][]int
	queue = append(queue, []int{0, 0})
	for len(graph) != 0{
		if len(queue) == 0{
			break
		}
		pop := queue[0]
		queue = queue[1:]
		for j := range graph[0]{
			queue = append(queue, []int{graph[0][j][1], graph[0][j][2] + pop[1]})
		}
		fmt.Println(queue)
		graph = graph[1:]
		queue = removeDuplicates(queue, pop)
		queue = sort(queue)
		res = checkValid(res, pop)
	}
	return res
}

func main(){
	n := 12
	edges := [][]int{
		{0, 1, 3},
		{0, 2, 4},
		{1, 2, 5}, 
		{1, 3, 4}, 
		{2, 4, 2}, 
		{3, 5, 3}, 
		{3, 6, 4}, 
		{4, 3, 5}, 
		{4, 5, 4}, 
		{4, 7, 4}, 
		{6, 8, 1}, 
		{7, 6, 5}, 
	}
	graph := make([][][]int, n)
	for _, v := range edges{
		graph[v[0]] = append(graph[v[0]], v)
	}
	fmt.Println(bfs_dijkstra(n, edges, graph))
}