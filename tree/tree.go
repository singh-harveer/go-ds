package tree

import (
	"fmt"
	"log"
	"math"
	"sync"
)

// Node represent Tree's node.
type Node struct {
	data  int
	left  *Node
	right *Node
}

type Binary struct {
	root *Node
	m    sync.RWMutex
}

// Insert add new value to binary tree.
func (b *Binary) Insert(value int) {
	b.m.Lock()
	defer b.m.Unlock()

	if b.root == nil {
		b.root = &Node{
			data: value,
		}

		return
	}

	insert(b.root, value)
}

// Height calculates height of tree.
func (b *Binary) Height() int {
	return height(b.root)

}

func height(root *Node) int {
	if root == nil {
		return 0 // height of empty tree.
	}

	var lHeight = height(root.left)
	var rHeight = height(root.right)
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

	var queue []*Node
	queue = append(queue, b.root)

	for len(queue) > 0 {
		var count = len(queue)

		for count > 0 {
			var currentNode = queue[0]
			queue = queue[1:] //remove currentNode from queue
			log.Print(currentNode.data)

			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}
			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}

			count--
		}
	}
}

// InOrderTraversal binary tree inorder traversal implementation.
func (b *Binary) InOrderTraversal() {
	inOrderTraversal(b.root)
}

// InOrderTraversal binary tree preorder traversal implementation.
func (b *Binary) PreOrderTraversal() {
	preOrderTraversal(b.root)
}

// InOrderTraversal binary tree PostOrder traversal implementation.
func (b *Binary) PostOrderTraversal() {
	postOrderTraversal(b.root)
}

// ConvertToMirrorImage convert given binary tree to it's mirror image.
// validate - in order traversal of mirror image whould reverse of normal tree.
func (b *Binary) ConvertToMirrorImage() {
	mirrorImageOfBinaryTree(b.root)
}

func mirrorImageOfBinaryTree(node *Node) {
	if node == nil {
		return
	}
	mirrorImageOfBinaryTree(node.left)
	mirrorImageOfBinaryTree(node.right)
	// swap left and right node
	var temp = node.left
	node.right = node.left
	node.left = temp
}

func postOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	postOrderTraversal(root.left)
	postOrderTraversal(root.right)
	fmt.Print(root.data)
}

func preOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data)
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}

func inOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	inOrderTraversal(root.left)
	fmt.Print(root.data)
	inOrderTraversal(root.right)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.data {
		return true
	} else if value > node.data {
		return search(node.right, value)
	} else {
		return search(node.left, value)
	}
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{
			data: value,
		}
	}

	if value > node.data {
		node.right = insert(node.right, value)
	} else {
		node.left = insert(node.left, value)
	}

	return node
}

func NewBinary() *Binary {
	return &Binary{}
}
