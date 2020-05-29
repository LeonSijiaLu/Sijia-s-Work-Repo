package main

import (
	"fmt"
)

func DiskStacking(disks [][]int) [][]int {
	if len(disks) == 1{return disks}

	queue := make([][]int, len(disks))
	for i := range disks{
		queue[i] = make([]int, len(disks[i]))
	}

	cur := disks[1]
	
}

func main(){
	disks := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{2, 2, 8}, 
		{2, 3, 4},
		{1, 3, 1},
		{4, 4, 5},
	}
	fmt.Println(DiskStacking(disks))
}
