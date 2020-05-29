package main


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
// next array 其实就是一个 matching pattern


func constructNext(substr string) []int{
	if len(substr) == 1{
		return []int{-1}
	}
	if len(substr) == 2{
		if substr[0] == substr[1]{
			return []int{-1, 0}
		}else{
			return []int{-1, -1}
		}
	}

	next := make([]int, len(substr))
	for i := range next{
		next[i] = -1
	}

	i, j := 1, 0
	for i < len(substr){
		if substr[i] == substr[j]{
			next[i] = j
			i, j = i + 1, j + 1
		}else if j > 0{
			j = next[j-1] + 1
		}else{
			i ++
		}
	}
	return next
}

func findDuplicates(str string, substr string, next []int) bool{
	i, j := 0, 0
	for i + len(substr) - j <= len(str) {
		if str[i] == substr[j]{
			if j == len(substr) - 1{
				return true
			}
			i, j = i + 1, j + 1
		}else if j > 0{
			j = next[j - 1] + 1
		}else{
			i ++
		}
	}
	return false
} 

func KnuthMorrisPrattAlgorithm(str, substr string) bool {
	next := constructNext(substr)
	return findDuplicates(str, substr, next)
}

func main(){
	str := "aefoaefcdaefcdaed"
	substr := "aefaedaefaefa"
	KnuthMorrisPrattAlgorithm(str, substr)
}