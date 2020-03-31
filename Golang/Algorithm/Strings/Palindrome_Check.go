package main

import (
	"fmt"
)

func IsEven (str string) bool{
	return len(str) % 2 == 0
}

func IsPalindrome(str string) bool{
	if(len(str) == 1 || len(str) == 0){return true}
	starting, ending, isEqual := 1, len(str), false
	isEven := IsEven(str)
	for (!isEqual){
		if(str[starting-1] == str[ending-1]){
			starting = starting +1
			ending = ending -1
			if isEven {
				if (starting == ending-1){isEqual = true}
			}else{
				if (starting == ending){isEqual = true}
			}
		}else{
			break
		}
	}
	if(isEqual == false){return false}else{return true}
}

func main(){
	fmt.Println(IsPalindrome("abcdefgfedcba"));
	fmt.Println(IsPalindrome("abba"));
	fmt.Println(IsPalindrome("abcba"));
	fmt.Println(IsPalindrome("abcdefgsdffedcba"));
}