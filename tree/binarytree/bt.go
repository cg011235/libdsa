package binarytree

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root      *Node
	nodeCount int
}

func (t *Tree) NumNodes() int {
	return t.nodeCount
}

func createHelper(arr []int, i int, n int) *Node {
	if i >= n {
		return nil
	}
	t := new(Node)
	t.value = arr[i]
	t.left = createHelper(arr, i+1, n)
	t.right = createHelper(arr, i+2, n)
	return t
}

func (t *Tree) Create(arr []int) {
	n := len(arr)
	if n <= 0 {
		return
	}
	t.root = createHelper(arr, 0, n)
	t.nodeCount = n
}

func copyTreeHelper(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	tmp := new(Node)
	tmp.value = curr.value
	tmp.left = copyTreeHelper(curr.left)
	tmp.right = copyTreeHelper(curr.right)
	return tmp
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTreeHelper(t.root)
	return Tree2
}

func isEqualHelper(node1 *Node, node2 *Node) bool {

	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) &&
			isEqualHelper(node1.left, node2.left) &&
			isEqualHelper(node1.right, node2.right))
	}
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqualHelper(t.root, t2.root)
}

func isCompleteTreeHelper(curr *Node, index int, count int) bool {
	if curr == nil {
		return true
	}
	if index >= count {
		return false
	}
	return isCompleteTreeHelper(curr.left, 2*index+1, count) &&
		isCompleteTreeHelper(curr.right, 2*index+2, count)
}

func (t *Tree) IsCompleteTreeRecursive() bool {
	count := t.NumNodes()
	return isCompleteTreeHelper(t.root, 0, count)
}
