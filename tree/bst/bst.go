package bst

import (
	"fmt"
	"math"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func CreateBinarySearchTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinarySearchTreeUtil(arr, 0, size-1)
	return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {
	if start <= end {
		curr := new(Node)
		mid := (start + end) / 2
		curr.value = arr[mid]
		curr.left = createBinarySearchTreeUtil(arr, start, mid-1)
		curr.right = createBinarySearchTreeUtil(arr, mid+1, end)
		return curr
	} else {
		return nil
	}
}

func (t *Tree) Add(value int) {
	t.root = addNode(t.root, value)
}

func addNode(n *Node, value int) *Node {
	if n == nil {
		return &Node{value, nil, nil}
	} else {
		if value < n.value {
			return addNode(n.left, value)
		} else {
			return addNode(n.right, value)
		}
	}
}

func (t *Tree) PrintInOrder() {
	if t.root != nil {
		printInOrder(t.root)

	} else {
		fmt.Print("Empty Tree")
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Print(n.value, " ")
	printInOrder(n.right)
}

func (t *Tree) Find(value int) bool {
	return traverseToFind(t.root, value)
}

func traverseToFind(n *Node, value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	return traverseToFind(n.left, value) || traverseToFind(n.right, value)
}

func (t *Tree) FindMinNode() *Node {
	var node *Node = t.root
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root
	for node != nil && node.right != nil {
		node = node.right
	}
	return node
}

func (t *Tree) IsBST() bool {
	return IsBST(t.root, math.MinInt32, math.MaxInt32)
}

func IsBST(curr *Node, min int, max int) bool {
	if curr == nil {
		return true
	}
	if curr.value < min || curr.value > max {
		return false
	}
	return IsBST(curr.left, min, curr.value) && IsBST(curr.right,
		curr.value, max)
}

func (t *Tree) DeleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			}
			if node.left == nil {
				temp = node.right
				return temp
			}
			if node.right == nil {
				temp = node.left
				return temp
			}
			maxNode := t.FindMax(node.left)
			maxValue := maxNode.value
			node.value = maxValue
			node.left = t.DeleteNode(node.left, maxValue)
		} else {
			if node.value > value {
				node.left = t.DeleteNode(node.left, value)
			} else {
				node.right = t.DeleteNode(node.right, value)
			}
		}
	}
	return node
}
