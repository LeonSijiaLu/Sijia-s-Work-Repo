package main

import (
	"fmt"
)

type Dijkstra_Node struct{
	visited bool
	value int
	path string
}

const INT_MAX = 1<<32 -1 

func Dijkstra_Shortest_Path(starting int, vertex [][]int) []Dijkstra_Node{
	if len(vertex) == 0 {return []Dijkstra_Node{}}
	finished := false

	result_q := make([]Dijkstra_Node, len(vertex)) // used for 
	result_q[0].visited = true
	result_q[0].value = 0

	for {
		if finished == true{
			break
		}
		for i:= 0; i < len(vertex); i ++{
			max := 0
			for j := range result_q{
				if result_q[j].visited == false{
					if max < result_q[j]{
						max = result_q[j]
					}
				}
			}
		}
	}

	for i := 0; i < len(vertex); i ++{
		min := INT_MAX
		min_index := INT_MAX
		for j := 0; j < len(vertex[i]); j ++{
			if vertex[i][j] < min{
				min = vertex[i][j]
				min_index = j
			}
		}
		
	}
}

func main(){
	var vertex [][]int // [0][2] 从 v0 到 v2 的距离
	for i := 0; i < 6; i ++{
		vertex = append(vertex, make([]int, 6))
	}

	for i := 0; i < 6; i++ {
		for j:=0;j <6; j++ {
			vertex[i][j] = INT_MAX
		}
	}

	starting := 2

	vertex[0][2] = 10
	vertex[0][4] = 30
	vertex[0][5] = 100
	vertex[1][2] = 5
	vertex[2][3] = 50
	vertex[3][5] = 10
	vertex[4][5] = 60
	vertex[4][3] = 20

	shortest_path := Dijkstra_Shortest_Path(starting, vertex)
	fmt.Println("Shortest Path", shortest_path)
}

