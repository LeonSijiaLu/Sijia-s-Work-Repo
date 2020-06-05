package main

import (
	"container/list"
	"fmt"
)

const MAX_INT = int(^uint(0) >> 1)

func bfs_dijkstra(n int, edges [][]int, src int, dst int, stops int) int{
	ret := MAX_INT // max distance
	graph := make([][][]int, n)
	for _, edge := range edges{
		graph[edge[0]] = append(graph[edge[0]], edge[1:])
	}
	// graphs[0]:
	// [0, 1, 200]
	// [0, 2, 200]
	// [0, 3, 100]

	// graphs[1]:
	// [1, 2, 500]

	// graphs[2]:
	// [2, 3, 50]
	queue := list.New()
	queue.PushBack([]int{src, 0})
	level := 0
	for queue.Len() != 0 && level <= stops + 1{
		size := queue.Len()
		for i := 0; i < size; i ++{ // 尝试把 queue 每个node都拿出来
			e := queue.Front()
			queue.Remove(e)
			cur := e.Value.([]int) // cur[0] is node, cur[1] is distance
			if cur[0] == dst && ret > cur[1] {ret = cur[1]} // if this node is dst and distance is the shortest
			for _, edge := range graph[cur[0]]{ // 找出这个node的所有边
				// cur[1] + edge[1] 这个node现在的距离加上这条边的距离
				if tempCost := cur[1] + edge[1]; tempCost < ret{
					queue.PushBack([]int{edge[0], tempCost})
				}
			}
		}
		level ++
	}
	if ret == MAX_INT{return -1}else{return ret}
}

func main(){
	n := 12
	edges := [][]int{
		{0, 1, 3},
		{0, 2, 4},
		{1, 2, 5}, 
		{1, 3, 4}, 
		{2, 4, 2}, 
		{4, 3, 5}, 
		{3, 5, 3}, 
		{3, 6, 4}, 
		{4, 5, 4}, 
		{4, 7, 4}, 
		{7, 6, 5}, 
		{6, 8, 1}, 
	}
	src := 0
	dst := 8
	k := 100
	fmt.Println(bfs_dijkstra(n, edges, src, dst, k))
}