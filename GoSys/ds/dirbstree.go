package ds

type Dir struct {
	FileTree FileBSTree
	DirTree  DirBSTree
	Name     string
}

type DirBSTreeNode struct {
	LeftNode  *DirBSTreeNode
	RightNode *DirBSTreeNode
	Value     Dir
}

type DirBSTree struct {
	Root *DirBSTreeNode
	Size int
}

func BuildDirNode() *DirBSTreeNode {
	return new(DirBSTreeNode)
}

func (node *DirBSTreeNode) Comp(other *DirBSTreeNode) bool {
	return node.Value.Name < other.Value.Name
}

func (tree *DirBSTree) Push(value Dir, comp func(Dir, Dir) bool) int {
	if (tree.Size == 0) || (tree.Root == nil) {
		tree.Root = BuildDirNode()
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
					now.LeftNode = BuildDirNode()
					now.LeftNode.Value = value
					break
				}
			} else {
				if now.RightNode != nil {
					now = now.RightNode
				} else {
					now.RightNode = BuildDirNode()
					now.RightNode.Value = value
					break
				}
			}
		}
		tree.Size += 1
		return tree.Size
	}
}

func (now *DirBSTreeNode) Traversal(res []DirBSTreeNode) []DirBSTreeNode {
	if now.LeftNode != nil {
		res = now.LeftNode.Traversal(res)
	}
	res = append(res, *now)
	if now.RightNode != nil {
		res = now.RightNode.Traversal(res)
	}
	return res
}

func (tree *DirBSTree) Find(target Dir, comp func(Dir, Dir) bool) bool {
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

func (tree *DirBSTree) Sort() []Dir {
	var res []Dir
	var rv []DirBSTreeNode
	rv = tree.Root.Traversal(rv)
	for _, node := range rv {
		res = append(res, node.Value)
	}
	return res
}
