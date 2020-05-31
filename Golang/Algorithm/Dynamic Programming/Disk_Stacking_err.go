package main

import (
	"fmt"
)

func getMax(a int, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}

func SortDisk(disks [][]int) [][]int{
	for i := 0; i < len(disks); i ++{
		for j := range disks{
			if j != len(disks)-1 && disks[j][2] > disks[j + 1][2]{
				disks[j], disks[j + 1] = disks[j + 1], disks[j]
			}
		}
	}
	return disks
}

func DiskStacking(disks [][]int) [][]int {
	disks = SortDisk(disks)

	if len(disks) == 1{return disks}
	queue := make([]int, len(disks))
	for i := range disks{
		queue[i] = disks[i][2]
	}

	var stacks [][]int
	var disk_index []int
	var max_disk_index []int
	max_height := 0

	for i := 1; i < len(disks); i ++{
		disk_index = disk_index[:0]
		height := disks[i][2]
		for j := 0; j < i; j ++{
			if disks[j][0] < disks[i][0] && disks[j][1] < disks[i][1] && disks[j][2] < disks[i][2]{
				temp := queue[i]
				queue[i] = getMax(queue[i], height + queue[j])
				if temp != queue[i]{
					disk_index = append(disk_index, j)
				}
			}
		}
		if max_height < queue[i]{
			max_height = queue[i]
			max_disk_index = max_disk_index[:0]
			max_disk_index = disk_index
			max_disk_index = append(max_disk_index, i)
		}
	}
	for _, v := range max_disk_index{
		stacks = append(stacks, disks[v])
	}
	fmt.Println(disks)
	fmt.Println(queue)
	fmt.Println(max_disk_index)
	return stacks
}

func main(){
	disks := [][]int{
		{3, 3, 4},
		{2, 1, 2},
		{3, 2, 3}, 
		{2, 2, 8},
		{2, 3, 4},
		{5, 5, 6},
		{1, 2, 1},
		{4, 4, 5},
		{1, 1, 4},
		{2, 2, 3},
	}
	fmt.Println(DiskStacking(disks))
}
