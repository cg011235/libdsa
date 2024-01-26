package bst

import "testing"

func TestBSTCreation(t *testing.T) {
	// Test with an empty array
	arr := make([]int, 0)
	bst := CreateBinarySearchTree(arr)
	if bst.root != nil {
		t.Errorf("Expected BST root to be nil for empty array")
	}

	// Test with an array containing one element
	arr = append(arr, 1)
	bst = CreateBinarySearchTree(arr)
	if bst == nil {
		t.Error("Expected allocated tree, got nil")
	} else if bst.root.value != 1 || bst.root.left != nil || bst.root.right != nil {
		t.Errorf("BST root not correctly created with array containing one element")
	}

	// Test creating a tree from a sorted array.
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bst = CreateBinarySearchTree(arr2)
	if bst == nil {
		t.Error("Expected allocated tree, got nil")
	} else if bst.root.value != 5 || bst.root.left.value != 2 || bst.root.right.value != 8 {
		t.Errorf("BST root not correctly created with sorted array")
	}
}

/**
Test Insertion:
Insert into an empty tree.
Insert into a non-empty tree.
Insert duplicate values (if your implementation allows or disallows duplicates).
Test Search Functionality:

Search for a value in an empty tree (should not be found).
Search for a value that exists in the tree.
Search for a value that does not exist in the tree.
Test Deletion:

Delete a leaf node.
Delete a node with one child.
Delete a node with two children.
Delete a node that does not exist in the tree.
Delete from an empty tree.
Test Tree Traversal:

In-order traversal of the tree and verify if the output is a sorted array.
Pre-order, Post-order, and Level-order traversals to verify they match the expected output.
Test Tree Properties:

Check if the left child of every node is smaller, and the right child is larger (for a BST with unique values).
Test the height of the tree and see if it is within expected limits (especially for randomly inserted values).
Test Edge Cases:

Insertion and deletion of values on the boundaries of the data type (like INT_MIN, INT_MAX for integers).
Handling of special cases like deletion of the root node.
Test Balancing (if implementing a self-balancing BST like AVL or Red-Black Tree):

Insert nodes in a way that triggers tree rotations and verify the balance of the tree after each insertion/deletion.
Test Memory Management:

If your language requires manual memory management (like C++), test for memory leaks or improper memory deallocation.
Randomized Tests:

Create a large number of random insertions and deletions and then verify the tree structure and properties.


**/
