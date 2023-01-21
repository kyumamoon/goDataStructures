package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type Tree struct {
	root *node
}

func (t *Tree) InsertNode(n *node) {
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

func SearchNode(current, parent *node, targetData int) (bool, *node, *node) {

	if targetData > current.data {
		if current.right != nil {
			return SearchNode(current.right, current, targetData)
		} else {
			return false, nil, nil
		}
	} else if targetData < current.data {
		if current.left != nil {
			return SearchNode(current.left, current, targetData)
		} else {
			return false, nil, nil
		}
	}

	return true, current, parent

}

func (t *Tree) SearchTree(targetData int) (bool, *node, *node) {
	if t.root == nil {
		return false, nil, nil
	}

	return SearchNode(t.root, nil, targetData)
}

func findChildlessChild(n *node) *node {
	if n.left == nil && n.right == nil {
		return n
	} else {
		return findChildlessChild(n.right)
	}
}

func (t *Tree) DeleteNode(targetNode, targetNodeParent *node) {
	// (PLAN)
	// First search left child node's right branch for child (with no children) with a value greater than its value.
	// Else, search right child node's left branch for child (with no children).

	// Check if target node is a root or child.
	if targetNodeParent == nil {
		// Deleting Root.

		// Check if root node has any children.

		// If both are branches are not empty.
		if targetNode.left != nil && targetNode.right != nil {

			childlessChild := findChildlessChild(targetNode.left)
			t.root = childlessChild // Replace deletion with childlessChild

		} else if targetNode.left == nil && targetNode.right == nil {
			t.root = nil
		} else {
			// Make either left or right branch node the new root.
			if targetNode.left != nil {
				t.root = targetNode.left
			} else {
				t.root = targetNode.right
			}
		}

	} else {

		// Deleting a child from a root.

		// Figure out whether target is on left or right side of parent.
		// 0 = left branch, 1 = right branch
		targetPosition := 0

		if targetNode.data > targetNodeParent.data {
			targetPosition = 1
		} else {
			targetPosition = 0
		}

		// Check if target has any children.

		// If target has two children.
		if targetNode.left != nil && targetNodeParent.right != nil {
			// Only search the left child's right branch.
			childlessChild := findChildlessChild(targetNode.left)

			// Replace the deleted node with the childLess child found above.
			// Find whether the deleted node was on left or right branch of parent node.
			// Then, replace the parent's branch with childlessChild.
			// Stitch the target's original left branch onto the childlessChild
			if targetPosition == 0 {
				targetNodeParent.left = childlessChild
				childlessChild.left = targetNode.left

			} else {
				targetNodeParent.right = childlessChild
				childlessChild.left = targetNode.left
			}

		} else if targetNode.left == nil && targetNode.right == nil {
			// If target has no children, replace parent's branch with nil.
			if targetPosition == 0 {
				targetNodeParent.left = nil
			} else {
				targetNodeParent.right = nil
			}
		} else {
			// Target has an empty left or right child
			// Make either left or right non-empty branch node the new target replacement.
			if targetNode.left != nil {
				if targetPosition == 0 {
					targetNodeParent.left = targetNode.left
				} else {
					targetNodeParent.right = targetNode.left
				}
			} else {
				if targetPosition == 0 {
					targetNodeParent.left = targetNode.right
				} else {
					targetNodeParent.right = targetNode.right
				}
			}
		}

	}
}

func main() {
	myTree := Tree{}
	x, y, z, w := &node{data: 1}, &node{data: 1}, &node{data: 2}, &node{data: 3}
	myTree.InsertNode(x)
	myTree.InsertNode(y)
	myTree.InsertNode(z)
	myTree.InsertNode(w)

	a, b, c := myTree.SearchTree(3)
	fmt.Println(a, b, c)
}
