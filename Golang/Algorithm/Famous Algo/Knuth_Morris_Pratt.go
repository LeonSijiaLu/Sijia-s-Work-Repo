package main

import (
	"fmt"
)

// string = "aefoaefcdaefcdaed"
// substring = "aefcdaed"

// 1. 前缀表，写出所偶遇前缀
// 								-1
// a							0
// ae							0
// aef							0
// aefc							0
// aefcd						0
// aefcda						0
// aefcdae  这是最长前缀		 2
// aefcdaed 这是本身, 不代入计算  0

// 2. 根据创建出来的前缀表，进行整理
// a	e	f	o	a	e	f	c	d	a	e	f	c	d	a	e	d 源码

// 0	1	2	3	4	5	6	7 序号
// a	e	f	c	d	a	e	d 符号
// -1	0	0	0	0	0	0	2 对应数字

// 3. 把 substring 和 string 对齐，一点点向后移动
// 如果匹配失败, 比如在 c 匹配失败的话
// c 对应数字是0
// 那么就把 序号 = 0 的和 o 比较

func constructNext(str string) int{
	counter := 0
	finished := false
	half_index := len(str) / 2
	left_arr := str[:half_index]
	right_arr := str[half_index:]
	if len(str) % 2 != 0{
		left_arr = str[:half_index]
		right_arr = str[half_index + 1:]
	}
	for {
		if finished == true{
			break
		}
		if len(left_arr) == 1 && len(right_arr) == 1{
			finished = true
		}
		for index, _ := range left_arr{
			if left_arr[index] == right_arr[index]{
				counter ++
			}else{
				left_arr = left_arr[:len(left_arr) - 1]
				right_arr = right_arr[1:]
				break
			}
			finished = true
		}
	}
	return counter
}

func compareStr(str string, substr string, next []int) int{
	starting := 0
	num := 0
	finished := false
	for {
		if finished == true{
			break
		}
		for i:=0; i<len(substr); i++{
			if substr[i] != str[starting]{
				if next[i] == -1{
					starting ++
					i = -1
				}else{
					i = next[i] - 1
				}
			}else{
				starting ++
			}
			if starting == len(str){
				finished = true
				break
			}
		}
		if finished == false{
			num ++
		}
	}
	return num
}

func KnuthMorrisPrattAlgorithm(str, substr string) bool {
	var next []int
	for index, _ := range substr{
		if index == 0{
			next = append(next, -1)
		}else if index == 1{
			if substr[1] == substr[0]{
				next = append(next, 1)
			}else{
				next = append(next, 0)
			}
		}else{
			next = append(next, constructNext(substr[:index]))
		}
	}
	fmt.Println(next)
	num := compareStr(str, substr, next)
	if num == 0{
		return false
	}else{
		return true
	}
}

func main(){
	str := "aefoaefcdaefcdaed"
	substr := "aefcdaed"
	res := KnuthMorrisPrattAlgorithm(str, substr)
	fmt.Println(res)
}