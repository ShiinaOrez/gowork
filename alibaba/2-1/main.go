package main

import "fmt"

type Link struct {
	Head *LinkNode
}

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func main() {
	link := Link{Head: &LinkNode{
		Value: 1,
		Next:  &LinkNode{
			Value: 2,
			Next:  &LinkNode{
				Value: 3,
				Next:  &LinkNode{
					Value: 4,
					Next:  &LinkNode{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}}
	link.printLink()
	LinkReverseK(&link, 5)
	link.printLink()
}

func LinkReverseK(link *Link, k int) {
	head := link.Head
	if k == 0 {
		return
	}
	linkLength := 0
	vHead := head
	var vTail *LinkNode
	for vHead != nil {
		linkLength += 1
		if vHead.Next == nil {
			vTail = vHead
		}
		vHead = vHead.Next
	}
	if k == linkLength {
		return
	}
	vHead = head
	for i:=0; i<linkLength-k-1; i++ {
		vHead = vHead.Next
	}
	(*link).Head = vHead.Next
	(*vTail).Next = head
	(*vHead).Next = nil
	return
}

func (head *Link) printLink() {
	start := head.Head
	for start != nil {
		fmt.Printf("%d ", start.Value)
		start = start.Next
	}
	fmt.Println()
}