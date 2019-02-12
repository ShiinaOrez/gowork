package ds

import (
	_ "fmt"
)

type QueueNode struct {
	Next   *QueueNode
	Value  Dir
}

type Queue struct {
	Head   *QueueNode
	Tail   *QueueNode
	Size   int
}

func (queue *Queue) Push (value Dir) int {
	newNode := new(QueueNode)
	newNode.Value = value
	if (queue.Head == nil) || (queue.Tail == nil) {
		queue.Head = newNode
		queue.Tail = newNode
		queue.Size = 1
	}else{
		queue.Tail.Next = newNode
		queue.Tail = newNode
		queue.Size += 1
	}
	return queue.Size
}

func (queue *Queue) Pop () {
	if (queue.Head == nil) || (queue.Size == 0) {
		return 
	}else{
		if (queue.Head == queue.Tail) || (queue.Size == 1) {
			queue.Head = nil
			queue.Tail = nil
			queue.Size = 0
			return
		}else{
			queue.Head = queue.Head.Next
			queue.Size -= 1
		}
	}
	return
}

func (queue *Queue) Peek () Dir{
	if (queue.Size > 0) && (queue.Head != nil) {
		return queue.Head.Value
	}else{
		return new(QueueNode).Value
	}
}

func (queue *Queue) Empty () bool {
	if (queue.Size == 0) || (queue.Head == nil) {
		return true
	}else{
		return false
	}
}