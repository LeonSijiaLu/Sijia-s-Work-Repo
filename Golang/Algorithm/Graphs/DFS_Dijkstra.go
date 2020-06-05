package main

import (
	"container/list"
	"fmt"
)

const MAX_INT = int(^uint(0) >> 1)

func bfs_dijkstra(n int, edges [][]int, src int, dst int, stops int) int{
	
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
	fmt.Println(dfs_dijkstra(n, edges, src, dst, k))
}