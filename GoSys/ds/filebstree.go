package ds

import (
	"github.com/ShiinaOrez/gowork/GoSys/typedefs"
)

type File typedefs.File

type FileBSTreeNode struct {
	LeftNode  *FileBSTreeNode
	RightNode *FileBSTreeNode
	Value     File
}

type FileBSTree struct {
	Root *FileBSTreeNode
	Size int
}

func BuildFileNode() *FileBSTreeNode {
	return new(FileBSTreeNode)
}

func (node *FileBSTreeNode) Comp(other *FileBSTreeNode) bool {
	return node.Value.Name < other.Value.Name
}

func (tree *FileBSTree) Push(value File, comp func(File, File) bool) int {
	if (tree.Size == 0) || (tree.Root == nil) {
		tree.Root = BuildFileNode()
		tree.Root.Value = value
		tree.Size = 1
		return tree.Size
	} else {
		now := tree.Root
		for now != nil {
			if comp(value, now.Value) {
				if now.LeftNode != nil {
					now = now.LeftNode
				} else {
					now.LeftNode = BuildFileNode()
					now.LeftNode.Value = value
					break
				}
			} else {
				if now.RightNode != nil {
					now = now.RightNode
				} else {
					now.RightNode = BuildFileNode()
					now.RightNode.Value = value
					break
				}
			}
		}
		tree.Size += 1
		return tree.Size
	}
}

func (now *FileBSTreeNode) Traversal(res []FileBSTreeNode) []FileBSTreeNode {
	if now.LeftNode != nil {
		res = now.LeftNode.Traversal(res)
	}
	res = append(res, *now)
	if now.RightNode != nil {
		res = now.RightNode.Traversal(res)
	}
	return res
}

func (tree *FileBSTree) Find(target File, comp func(File, File) bool) bool {
	now := tree.Root
	for now != nil {
		if now.Value == target {
			return true
		} else {
			if comp(target, now.Value) {
				now = now.LeftNode
			} else {
				now = now.RightNode
			}
		}
	}
	return false
}

func (tree *FileBSTree) Sort() []File {
	var res []File
	var rv []FileBSTreeNode
	rv = tree.Root.Traversal(rv)
	for _, node := range rv {
		res = append(res, node.Value)
	}
	return res
}
