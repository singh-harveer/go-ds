package treeutils

import (
	"fmt"
	"golang/ds/tree"
	"math"
)

// This Package will contains all traversal
// implementations.

/*      	1
		   / \
		  /   \
	      2    3
		 / \   / \
        12  11 6  4
*/

// PreOrder process tree node using Preorder traversal.
// O/P : 1,2,12,11,3,6,4
func PreOrder(node *tree.Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// InOrder process tree node using InOrder traversal.
// O/P: 12,2,11,1,6,3,4
func InOrder(node *tree.Node) {
	if node == nil {
		return
	}

	InOrder(node.Left)
	fmt.Printf("%d ", node.Data)
	InOrder(node.Right)
}

// PostOrder process tree node using PostOrder traversal.
// O/P : 12,11,2,6,4,3,1
func PostOrder(node *tree.Node) {
	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Printf("%d ", node.Data)
}

// LevelOrder treverse tree in level by level
// this is also known as BFS.
func LevelOrder(node *tree.Node) {
	if node == nil {
		return
	}

	var queue []*tree.Node
	queue = append(queue, node)

	for len(queue) > 0 {
		var count = len(queue) //to process all nodes of same level at once.

		for count > 0 {
			var currentNode = queue[0] // dequeue
			queue = queue[1:]          // removed node after dequeue

			fmt.Printf("%d ", currentNode.Data)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

			count--
		}

		fmt.Println() // break line after every level.
	}
}

// Height find height of tree.
func Height(node *tree.Node) int {
	if node == nil {
		return 0
	}

	var lHeight = Height(node.Left)
	var rHeight = Height(node.Right)
	var height = math.Max(float64(lHeight), float64(rHeight)) + 1

	return int(height)
}

// CountNodeRecursive return nodes count of binary tree
func CountNodeRecursive(node *tree.Node) int {
	if node == nil {
		return 0
	}
	var lNodes = Height(node.Left)
	var rNodes = +Height(node.Right)

	return lNodes + rNodes + 1
}

// NodesCountIterative return nodes count of binary tree
func NodesCountIterative(node *tree.Node) int {
	if node == nil {
		return 0
	}

	var nodesCount int
	var queue []*tree.Node

	queue = append(queue, node)

	for len(queue) > 0 {

		var current = queue[0]
		queue = queue[1:]

		nodesCount++ // increament node count

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return nodesCount
}
