package main

import (
	"fmt"
)

func getMin(a int, b int, c int) int{
	if a < b{
		if a < c{
			return a
		}else{
			if b < c{
				return b
			}else{
				return c
			}
		}
	}else{
		if b < c{
			return b
		}else{
			if a < c{
				return a
			}else{
				return c
			}
		}
	}
}

func LevenshteinDistance(a, b string) int {
	if a == b{return 0}
	a = " " + a
	b = " " + b
	matrix := make([][]int, len(a))
	
	for i := range a{
		matrix[i] = make([]int, len(b))
		matrix[i][0] = i
	}
	for i := range b{
		matrix[0][i] = i
	}
	for i := 1; i < len(a); i ++{
		for j := 1; j < len(b); j ++{
			if a[i] == b[j]{
				matrix[i][j] = matrix[i-1][j-1]
			}else{
				matrix[i][j] = getMin(matrix[i - 1][j], matrix[i][j - 1], matrix[i - 1][j - 1]) + 1
			}
		}
	}
	return matrix[len(a) - 1][len(b) - 1]
}

func main(){
	str1 := "abc"
	str2 := "yabd"
	fmt.Println(LevenshteinDistance(str1, str2))
}