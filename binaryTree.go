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

func SearchNode(n *node, target int) bool {
	current := n

	if target == current.data {
		return true
	} else if target > current.data {
		if current.right != nil {
			current = current.right
			return SearchNode(current, target)
		} else {
			return false
		}
	} else {
		if current.left != nil {
			current = current.left
			return SearchNode(current, target)
		} else {
			return false
		}
	}
}

func (t *Tree) SearchTree(target int) bool {
	if t.root == nil {
		return false
	}
	current := t.root

	return SearchNode(current, target)
}

func main() {
	myTree := Tree{}
	node1 := node{data: 100}
	node2 := node{data: 200}
	myTree.Insert(&node1)
	myTree.Insert(&node2)
	fmt.Println(myTree.root)
	fmt.Println(myTree.root.right)

	fmt.Println(myTree.SearchTree(100))
}
