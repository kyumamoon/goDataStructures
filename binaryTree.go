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

func main() {
	myTree := Tree{}
	node1 := node{data: 100}
	node2 := node{data: 200}
	myTree.Insert(&node1)
	myTree.Insert(&node2)
	fmt.Println(myTree.root)
	fmt.Println(myTree.root.right)
}
