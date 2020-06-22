package main

import (
	"fmt"
	"sort"
)

func GroupAnagrams(words []string) [][]string {
	if len(words) == 0{return [][]string{}}
	if len(words) == 1{return [][]string{words}}

	checked := []int{}
	for i := 0; i < len(words); i ++{
		checked = append(checked, 0) 
	}

	var sorted_Words []string
	for i := range words{
		sorted_Words = append(sorted_Words, toSortString(words[i]))
	}

	var anagrams [][]string
	for i := range sorted_Words{
		anagram := []string{}
		for j := i + 1; j < len(sorted_Words); j ++{
			if sorted_Words[i] == sorted_Words[j]{
				if checked[j] == 0{
					anagram = append(anagram, words[j])
					checked[j] = 1
				}
			}
		}
		if checked[i] == 0{
			anagram = append(anagram, words[i])
			checked[i] = 1
		}
		if len(anagram) != 0{
			anagrams = append(anagrams, anagram)
		}
	}

	return anagrams
}

func toSortString(str string) string{
	wordBytes := []byte(str)
	sort.Slice(wordBytes, func(i int, j int) bool{
		return wordBytes[i] < wordBytes[j]
	})
	return string(wordBytes)
}

func main(){
	words := []string{
		"yo",
		"act",
		"flop",
		"tac",
		"cat",
		"oy",
		"olfp",
	}
	fmt.Println(GroupAnagrams(words))
}