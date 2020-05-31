package main

import (
	"fmt"
)

func getColumn(chess [8]int, row int) int{
	return chess[row]
}

func queen(chess [8]int, cur_row int){
	if cur_row == len(chess){
		fmt.Println(chess)
		return
	}
	for col := range chess{
		chess[cur_row] = col 
		flag := true
		for j := 0; j < cur_row; j ++{
			dis := col - getColumn(chess, j) // Big Column - Small Column
			if dis < 0 {dis = -dis}
			if col == getColumn(chess, j) || dis == cur_row - j{
				flag = false
				break
			}
		}
		if flag{
			queen(chess, cur_row + 1)
		}
	}
}

func main(){
	chess := [8]int{0,0,0,0,0,0,0,0}
	queen(chess, 0)
}