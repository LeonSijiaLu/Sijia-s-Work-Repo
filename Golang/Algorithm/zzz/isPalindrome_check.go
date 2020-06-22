package main

import (
	"fmt"
)

func isPalindrome(word string) bool{
	if len(word) == 1 || len(word) == 0{return true}
	start, end, finished := 0, len(word) - 1, false
	isEven := false
	if len(word) % 2 == 0{
		isEven = true
	}

	for !finished{
		if word[start] == word[end]{
			if isEven{
				if start == end - 1{
					return true
				}
			}else{
				if start == end{
					return true
				}
			}
			start ++
			end --
		}else{
			return false
		}
	}
	return false
}

func main(){
	str := "abcdcba"
	fmt.Println(isPalindrome(str))
}