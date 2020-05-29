package main

import (
	"fmt"
	"os"
)

type point struct{
	i, j int
}

var dirs = [4]int{{-1, 0}, {1, 0}, {0, 1}, (0, -1)}

var maze = {
	{0, 1, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{1, 1, 1, 0, 0},
	{0, 1, 0, 0, 1},
	{0, 1, 0, 0, 0},
}

func (*p point) add(r point) point{
	return &point{i: p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	//检查行是否越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	//检查列是否越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int{
	steps := make([][]int, len(maze))
	for i := range steps{
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start} // store poped out queue
	for len(Q) > 0{
		cur := Q[0] 
		Q= Q[1:]
		if cur == end{
			break
		}
		for _, dir := range dirs{
			next := cur.add(dir)
			val, ok := next.at(maze)
			if !ok || val != 0 {
				continue
			}
			if next == start{
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] == curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func getSteps(maze [][]int, end point){
	return maze[end.i][end.j]
}

func main(){
	steps := walk(maze, &{i: 0, j: 0}, &{i: len(maze) - 1, j: len(maze[0]) - 1})
	fmt.Println(steps)
	num := getSteps(steps, steps[len(maze) - 1][len(maze)[0] - 1])
}