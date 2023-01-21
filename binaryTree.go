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

func (t *Tree) DeleteNode(target int) {

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

	fmt.Println(myTree.SearchTree(0))
}
