package main

import (
	"fmt"
)

type Point struct{
	x int
	y int
}

var dirs = [4]Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1},}

func inRange(matrix[][]int, node Point) bool{
	if node.y >= len(matrix) || node.y < 0{
		return false
	}
	if node.x >= len(matrix[0]) || node.x < 0{
		return false
	}
	return true
}

func getNode(matrix [][]int, node Point, dir Point) (Point, bool){
	new_node := Point{x: node.x + dir.x, y: node.y + dir.y}
	if inRange(matrix, new_node){
		return new_node, true
	}else{
		return new_node, false
	}
}

func RiverSizes(matrix [][]int) []int {
	var result []int
	visited := make([][]bool, len(matrix))
	for i := range matrix{
		visited[i] = make([]bool, len(matrix[i]))
	}
	for i := range matrix{
		for j := range matrix[i]{
			visited[i][j] = false
		}
	}

	for i := range matrix{
		for j := range matrix[i]{
			if visited[i][j] == false{ // NOT visited
				res := traverse(i, j, visited, matrix)
				fmt.Println(res)
				if res != 0{
					result = append(result, res)
				}
			}
		}
	}
	return result
}

func traverse(i int, j int, visited [][]bool, matrix [][]int) int{
	if matrix[i][j] == 0{
		visited[i][j] = true
		return 0
	}
	length := 0
	counter := 0
	finished := false
	i_temp := i
	j_temp := j
	for !finished{
		for _, dir := range dirs{
			counter = 0
			node, ok := getNode(matrix, Point{x: j_temp, y: i_temp}, dir)
			if ok == true{
				if matrix[node.y][node.x] == 1{
					length ++
					counter ++
					j_temp = node.x
					i_temp = node.y
					visited[i_temp][j_temp] = true
				}else{
					j_temp = node.x
					i_temp = node.y
					visited[i_temp][j_temp] = true
				}
			}
		}
		if counter == 0{
			finished = true
		}
	}
	visited[i][j] = true
	return length
}

func main(){
	river := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(RiverSizes(river))
}