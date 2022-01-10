package bst

import (
	"fmt"
	"golang/ds/tree"
	"log"
	"math"
	"sync"
)

type Binary struct {
	root *tree.Node
	m    sync.RWMutex
}

// Insert add new value to binary tree.
func (b *Binary) Insert(value int) *tree.Node {
	b.m.Lock()
	defer b.m.Unlock()

	if b.root == nil {
		b.root = &tree.Node{
			Data: value,
		}

		return b.root
	}

	return insert(b.root, value)
}

// Height calculates height of tree.
func (b *Binary) Height() int {
	return height(b.root)

}

func height(root *tree.Node) int {
	if root == nil {
		return 0 // height of empty tree.
	}

	var lHeight = height(root.Left)
	var rHeight = height(root.Right)
	return int(math.Max(float64(lHeight), float64(rHeight))) + 1

}

// Search search value in binary tree.
func (b *Binary) Search(value int) bool {
	return search(b.root, value)
}

// LevelOrderTraversal traverse tree in BFS manner.
func (b *Binary) LevelOrderTraversal() {
	if b.root == nil {
		return
	}

	var queue []*tree.Node
	queue = append(queue, b.root)

	for len(queue) > 0 {
		var count = len(queue)

		for count > 0 {
			var currentNode = queue[0]
			queue = queue[1:] //remove currentNode from queue
			log.Print(currentNode.Data)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

			count--
		}
	}
}

// ConvertToMirrorImage convert given binary tree to it's mirror image.
// validate - in order traversal of mirror image whould reverse of normal tree.
func (b *Binary) ConvertToMirrorImage() {
	mirrorImageOfBinaryTree(b.root)
}

func mirrorImageOfBinaryTree(node *tree.Node) {
	if node == nil {
		return
	}
	mirrorImageOfBinaryTree(node.Left)
	mirrorImageOfBinaryTree(node.Right)
	// swap left and right node
	var temp = node.Left
	node.Right = node.Left
	node.Left = temp
}

func postOrderTraversal(root *tree.Node) {
	if root == nil {
		return
	}

	postOrderTraversal(root.Left)
	postOrderTraversal(root.Right)
	fmt.Print(root.Data)
}

func search(node *tree.Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Data {
		return true
	} else if value > node.Data {
		return search(node.Right, value)
	} else {
		return search(node.Left, value)
	}
}

func insert(node *tree.Node, value int) *tree.Node {
	if node == nil {
		return &tree.Node{
			Data: value,
		}
	}

	if value > node.Data {
		node.Right = insert(node.Right, value)
	} else {
		node.Left = insert(node.Left, value)
	}

	return node
}

func NewBinary() *Binary {
	return &Binary{}
}
