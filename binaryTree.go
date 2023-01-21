package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type Tree struct {
	root *node
}

// Insert inserts a node into a binary tree.
func (t *Tree) Insert(n *node) {
	if t.root != nil {
		current := t.root

		for {
			if n.data > current.data {
				if current.right != nil {
					current = current.right
				} else {
					current.right = n
					break
				}
			} else {
				if current.left != nil {
					current = current.left
				} else {
					current.left = n
					break
				}
			}
		}
	} else {
		t.root = n
	}
}

// SearchNode searches the node branch for the target value and returns true/false if found and pointer to the node containing value and a pointer to the node's parent.
func SearchNode(current, parent *node, target int) (bool, *node, *node) {

	if target > current.data {
		if current.right != nil {
			return SearchNode(current.right, current, target)
		} else {
			return false, nil, nil
		}
	} else if target < current.data {
		if current.left != nil {
			return SearchNode(current.left, current, target)
		} else {
			return false, nil, nil
		}
	}

	return true, current, parent

}

// SearchTree searches for the node that has the target data, returns bool if found/not found and the node and its parent if true.
func (t *Tree) SearchTree(target int) (bool, *node, *node) {
	if t.root == nil {
		return false, nil, nil
	}
	current := t.root

	return SearchNode(current, nil, target)
}

func SearchNodeBranchForEmpty(n *node) *node {
	if n.left == nil && n.right == nil {
		return n
	} else {

		return SearchNodeBranchForEmpty(n.right)

	}
}

// DeleteNode deletes a node from the binary tree and restitch the children to correct order.
func (t *Tree) DeleteNode(target, targetParent *node) {
	// (PLAN)
	// First search left child node's right branch for child (with no children) with a value greater than its value.
	// Else, search right child node's left branch for child (with no children).

	// Check if target node is a root or child.
	if targetParent == nil {
		// Deleting Root.

		// Check if root node has any children.

		// If both are branches are not empty.
		if target.left != nil && target.right != nil {

			current := target.left
			childlessChild := SearchNodeBranchForEmpty(current)
			t.root = childlessChild // Replace deletion with childlessChild

		} else {
			// Make either left or right branch node the new root.
			if target.left != nil {
				t.root = target.left
			} else {
				t.root = target.right
			}
		}

	} else {

		// Deleting a child from a root.

		// Figure out whether target is on left or right side of parent.
		targetLeftRight := 0

		if target.data > targetParent.data {
			targetLeftRight = 1
		} else {
			targetLeftRight = 0
		}

		// Check if target has any children.

		// If target has two children.
		if target.left != nil && target.right != nil {
			// Only search the left child's right branch.
			current := target.left

			childlessChild := SearchNodeBranchForEmpty(current)

			// Replace the deleted node with the childLess child found above.
			// Find whether the deleted node was on left or right branch of parent node.
			// Then, replace the parent's branch with childlessChild.
			// Stitch the target's original left branch onto the childlessChild
			if targetLeftRight == 0 {
				targetParent.left = childlessChild
				childlessChild.left = target.left

			} else {
				targetParent.right = childlessChild
				childlessChild.left = target.left
			}

		} else {
			// Target has an empty left or right child
			// Make either left or right non-empty branch node the new target replacement.
			if target.left != nil {
				if targetLeftRight == 0 {
					targetParent.left = target.left
				} else {
					targetParent.right = target.left
				}
			} else {
				if targetLeftRight == 0 {
					targetParent.left = target.right
				} else {
					targetParent.right = target.right
				}
			}
		}

	}
}

func main() {
	myTree := Tree{}
	myTree.Insert(&node{data: 44})
	myTree.Insert(&node{data: 17})
	myTree.Insert(&node{data: 32})
	myTree.Insert(&node{data: 28})
	myTree.Insert(&node{data: 29})
	myTree.Insert(&node{data: 88})
	myTree.Insert(&node{data: 65})
	myTree.Insert(&node{data: 54})
	myTree.Insert(&node{data: 82})
	myTree.Insert(&node{data: 76})
	myTree.Insert(&node{data: 80})
	myTree.Insert(&node{data: 78})
	myTree.Insert(&node{data: 97})

	x, y, z := myTree.SearchTree(32)
	fmt.Println(x, y, z)
	myTree.DeleteNode(y, z)

	x, y, z = myTree.SearchTree(28)
	fmt.Println(x, y, z)
}
