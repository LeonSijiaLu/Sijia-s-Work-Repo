package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func FindLoop(head *LinkedList) *LinkedList {
	p := head
	hash := make(map[int]int)
	for p != nil{
		_, ok := hash[p.Value]
		if ok { // exists
			return p
		}else{
			hash[p.Value] = p.Value
			p = p.Next
		}
	}
	return head
}

func (l *LinkedList) TraverseList(){
	fmt.Println(l.Value)
	if l.Next != nil{
		l.Next.TraverseList()
	}
}

func main(){
	l0 := &LinkedList{Value: 0, Next: nil,}
	l1 := &LinkedList{Value: 1, Next: nil,}
	l2 := &LinkedList{Value: 2, Next: nil,}
	l3 := &LinkedList{Value: 3, Next: nil,}
	l4 := &LinkedList{Value: 4, Next: nil,}
	l5 := &LinkedList{Value: 5, Next: nil,}
	l6 := &LinkedList{Value: 6, Next: nil,}
	l7 := &LinkedList{Value: 7, Next: nil,}
	l8 := &LinkedList{Value: 8, Next: nil,}
	l9 := &LinkedList{Value: 9, Next: nil,}
	
	l0.Next = l1
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l7.Next = l8
	l8.Next = l9
	l9.Next = l4

	loop := FindLoop(l0)
	loop.TraverseList()
}