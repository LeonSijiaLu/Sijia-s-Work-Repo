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

func findShortest(cur []int, steps [][]int, maze[][]int, dirs [][]int, width, length) []int{
	min := 100
	var min_dir []int
	for i := range dirs{
		if checkRange(cur[1] + dirs[i][1], cur[0] + dirs[i][0], width, length){	
			if maze[cur[1] + dirs[i][1]][cur[0] + dirs[i][0]] == 0{
				if steps[cur[1] + dirs[i][1]][cur[0] + dirs[i][0]] < min{
					min = steps[cur[1] + dirs[i][1]][cur[0] + dirs[i][0]]
					min_dir = min_dir[0:0]
					min_dir = append(min_dir, dirs[i][0])
					min_dir = append(min_dir, dirs[i][1])
				}
			}
		}
	}
	return min_dir
}

func run_dfs(cur []int, ending []int, steps [][]int, maze [][]int, dirs [][]int, width int, length int, Q [][]int){
	finished := false
	out_counter := 0
	wall_counter := 0
	for !finished{
		fmt.Println(cur, findShortest(cur, steps, maze, dirs, width, length))
		out_counter = 0
		wall_counter = 0
		for i := range dirs{
			if checkRange(cur[1] + dirs[i][1], cur[0] + dirs[i][0], width, length){	
				if cur[0] + dirs[i][0] == ending[0] && cur[1] + dirs[i][1] == ending[1]{
					finished = true
					break
				}
				if maze[cur[1] + dirs[i][1]][cur[0] + dirs[i][0]] == 0{
					Q = append(Q, []int{cur[0] + dirs[i][0], cur[1] + dirs[i][1]})
					steps[cur[1] + dirs[i][1]][cur[0] + dirs[i][0]] = steps[cur[1] + dirs[i][1]][cur[0] + dirs[i][0]] + 1
					cur = []int{cur[0] + dirs[i][0], cur[1] + dirs[i][1]}
					break
				}else{
					wall_counter ++
				}
			}else{
				out_counter ++
			}
		}
		if wall_counter == 3{
			cur = Q[len(Q) - 1]
			Q = Q[:len(Q) - 1]
		}
		if out_counter == 3{
			cur = Q[len(Q) - 1]
			Q = Q[:len(Q) - 1]
		}
	}
}

func dfs_maze(maze [][]int, width int, length int, dirs [][]int, starting []int, ending []int){
	var Q [][]int
	Q = append(Q, starting)
	steps := make([][]int, len(maze))
	for i := range maze{
		steps[i] = make([]int, len(maze[i]))
	}

	cur := starting

	run_dfs(cur, ending, steps, maze, dirs, width, length, Q)
	fmt.Println(Q)
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
	dfs_maze(maze, width, length, dirs, starting, ending)
}