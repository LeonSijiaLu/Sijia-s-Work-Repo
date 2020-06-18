package main

import (
	"fmt"
)

func maxMazeSize(maze [][]int) (int, int){
	return len(maze[0]), len(maze)
}

func checkRange(x int, y int, width int, length int) bool{
	if x >= width || y >= length || x < 0 || y < 0{
		return false
	}
	return true
}

func bfs_maze(maze [][]int, dirs [][]int, start []int, dst []int, width int, length int) [][]int{
	var Q [][]int
	Q = append(Q, start)
	steps := make([][]int, len(maze))
	for i := range maze{
		steps[i] = make([]int, len(maze[i]))
	}

	for len(Q) > 0{
		cur := Q[0]
		Q= Q[1:]
		if cur[0] == dst[0] && cur[1] == dst[1]{
			break
		}
		for i := range dirs{
			next := []int{cur[0] + dirs[i][0], cur[1] + dirs[i][1]}
			if next[0] == dst[0] && next[1] == dst[1]{
				steps[next[1]][next[0]] = steps[next[1]][next[0]] + 1
				return steps
			}
			if checkRange(next[1], next[0], width, length) == false{continue} // with in range
			if maze[next[1]][next[0]] == 1{continue} // not wall
			if next[0] == start[0] && next[1] == start[1]{continue} // not loop
			steps[next[1]][next[0]] = steps[next[1]][next[0]] + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main(){
	maze := [][]int{        // [y][x]
		{1, 0, 1, 0, 0, 0}, // length [0] = y
		{0, 0, 0, 0, 1, 0}, // length 
		{0, 0, 1, 0, 0, 0}, // length
		{0, 1, 1, 1, 0, 0}, // length
		{0, 0, 1, 0, 0, 1}, // length
		{1, 0, 1, 0, 0, 0}, // length
		// width, width [1] = x
	} 
	dirs := [][]int{
		{1, 0}, // 右
		{0, -1}, // 下
		{-1, 0}, // 左
		{0, 1}, // 上
	}
	starting := []int{0, 1} // y [0], x [1]
	ending := []int{5, 4}

	width, length := maxMazeSize(maze)
	res := bfs_maze(maze, dirs, starting, ending, width, length)
	fmt.Println(res[ending[1]][ending[0]])
}