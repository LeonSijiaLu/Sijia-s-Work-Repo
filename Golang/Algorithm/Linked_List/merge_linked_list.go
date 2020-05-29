package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	p1 := headOne
	p2 := headTwo
	finished := false
	for !finished{
		if p2 == headTwo && p1.Value <= p2.Value{
			headOne = headOne.Next
			p1.Next = headTwo
			headTwo = p1
			p1 = headOne
			p2 = headTwo
		}else if p2.Next != nil{
			if p2.Value <= p1.Value && p2.Next.Value >= p1.Value{
				headOne = p1.Next
				p1.Next = p2.Next
				p2.Next = p1
				p1 = headOne
				p2 = headTwo
			}else{
				p2 = p2.Next
			}
		}else if p2.Next == nil && p2.Value <= p1.Value{
			headOne = p1.Next
			p1.Next = p2.Next
			p2.Next = p1
			p1 = headOne
			p2 = headTwo
		}
		if headOne == nil{
			finished = true
		}
	}
	return headTwo
}

func (l *LinkedList) TraverseList(){
	fmt.Println(l.Value)
	if l.Next != nil{
		l.Next.TraverseList()
	}
}

func main(){
	l1 := &LinkedList{Value: 2, Next: nil,}
	l2 := &LinkedList{Value: 6, Next: nil,}
	l3 := &LinkedList{Value: 7, Next: nil,}
	l4 := &LinkedList{Value: 8, Next: nil,}

	m1 := &LinkedList{Value: 1, Next: nil,}
	m2 := &LinkedList{Value: 3, Next: nil,}
	m3 := &LinkedList{Value: 4, Next: nil,}
	m4 := &LinkedList{Value: 5, Next: nil,}
	m5 := &LinkedList{Value: 9, Next: nil,}
	m6 := &LinkedList{Value: 10, Next: nil,}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	m1.Next = m2
	m2.Next = m3
	m3.Next = m4
	m4.Next = m5
	m5.Next = m6

	mergedList := MergeLinkedLists(l1, m1)
	mergedList.TraverseList()
}