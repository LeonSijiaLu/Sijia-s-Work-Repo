// Change 1 2 3 4 5 6 7 8 9
// To 9 8 7 6 5 4 3 2 1

package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	p1 := head
	if p1.Next == nil{
		return p1
	}

	p2 := head.Next
	if p2.Next == nil{
		p1.Next = nil
		p2.Next = p1
		return p2
	}

	p3 := head.Next.Next
	if p3.Next == nil{
		p1.Next = nil
		p2.Next = p1
		p3.Next = p2
		return p3
	}

	finished := false
	p1.Next = nil
	for {
		p2.Next = p1
		if finished == false{
			p1 = p2
			p2 = p3
			p3 = p3.Next
		}else{
			p3.Next = p2
			break
		}
		if p3.Next == nil{
			finished = true
		}
	}
	return p3
}

func (l *LinkedList) TraverseList(){
	fmt.Println(l.Value)
	if l.Next != nil{
		l.Next.TraverseList()
	}
}

func main(){
	l1 := &LinkedList{Value: 1, Next: nil,}
	l2 := &LinkedList{Value: 2, Next: nil,}
	l3 := &LinkedList{Value: 3, Next: nil,}
	l4 := &LinkedList{Value: 4, Next: nil,}
	l5 := &LinkedList{Value: 5, Next: nil,}
	l6 := &LinkedList{Value: 6, Next: nil,}
	l7 := &LinkedList{Value: 7, Next: nil,}
	l8 := &LinkedList{Value: 8, Next: nil,}
	l9 := &LinkedList{Value: 9, Next: nil,}
	l10 := &LinkedList{Value: 10, Next: nil,}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l7.Next = l8
	l8.Next = l9
	l9.Next = l10

	ReverseLinkedList(l1)
	l10.TraverseList()
}