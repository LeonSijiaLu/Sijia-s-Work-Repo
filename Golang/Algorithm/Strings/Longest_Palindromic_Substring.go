package main

import (
	"fmt"
)

func longest_palindromic_substring(str string) string{
	longest_sub := ""
	for i := range str{
		for j := len(str) - 1; j >= 0; j --{
			if j < i {break}
			if str[j] == str[i]{
				res := isPanlindromic(str[i: j + 1])
				if len(res) > len(longest_sub){
					longest_sub = res
				}
			}
		}
	}
	return longest_sub
}

func isPanlindromic(str string) string{
	if len(str) == 0 || len(str) == 1{return str}
	start, end, finished := 0, len(str) - 1, false
	isEven := false
	if len(str) % 2 == 0{isEven = true}
	for !finished{
		if str[start] == str[end]{
			if isEven{
				if start == end - 1{
					return str
				}
			}else{
				if start == end{
					return str
				}
			}
			start ++
			end --
		}else{
			return ""
		}
	}
	return ""
}

func main(){
	str := "abaxyzzyxf"
	fmt.Println(longest_palindromic_substring(str))
}