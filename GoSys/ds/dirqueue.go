package ds

import (
	_ "fmt"
)

type PathNode struct {
	Next   *PathNode
	Value  Dir
}

type Path struct {
	Head   *PathNode
	Tail   *PathNode
	Size   int
}

func (path *Path) Push (value Dir) int {
	newNode := new(PathNode)
	newNode.Value = value
	if (path.Head == nil) || (path.Tail == nil) {
		path.Head = newNode
		path.Tail = newNode
		path.Size = 1
	}else{
		path.Tail.Next = newNode
		path.Tail = newNode
		path.Size += 1
	}
	return path.Size
}

func (path *Path) Pop () {
	if (path.Head == nil) || (path.Size == 0) {
		return 
	}else{
		if (path.Head == path.Tail) || (path.Size == 1) {
			path.Head = nil
			path.Tail = nil
			path.Size = 0
			return
		}else{
			path.Head = path.Head.Next
			path.Size -= 1
		}
	}
	return
}

func (path *Path) Peek () Dir{
	if (path.Size > 0) && (path.Head != nil) {
		return path.Head.Value
	}else{
		return new(PathNode).Value
	}
}

func (path *Path) Empty () bool {
	if (path.Size == 0) || (path.Head == nil) {
		return true
	}else{
		return false
	}
}
