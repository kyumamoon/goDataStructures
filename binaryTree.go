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

	if target > n.data {
		if n.right != nil {
			return SearchNode(n.right, target)
		} else {
			return false
		}
	} else if target < n.data {
		if n.left != nil {
			return SearchNode(n.left, target)
		} else {
			return false
		}
	}

	return true

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
	myTree.Insert(&node{data: 100})
	myTree.Insert(&node{data: 52})
	myTree.Insert(&node{data: 203})
	myTree.Insert(&node{data: 19})
	myTree.Insert(&node{data: 76})
	myTree.Insert(&node{data: 150})
	myTree.Insert(&node{data: 310})
	myTree.Insert(&node{data: 7})
	myTree.Insert(&node{data: 24})
	myTree.Insert(&node{data: 88})
	myTree.Insert(&node{data: 276})

	fmt.Println(myTree.SearchTree(76))
}
