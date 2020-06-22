package main

import (
	"fmt"
)

func LongestSubstringWithoutDuplication(str string) string {
	last_visited := map[string]int{}
	res := ""
	temp_res := ""
	for i := range str{
		val, ok := last_visited[string(str[i])]
		if !ok{
			temp_res = temp_res + string(str[i])
		}else{
			if len(temp_res) > len(res){
				res = temp_res
				temp_res = ""
			}
		}
	}
}

func main(){
	str := "clementisacap"
	fmt.Println(LongestSubstringWithoutDuplication(str))
}