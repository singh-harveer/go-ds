package main

import (
	"fmt"
	"golang/ds/tree"
	"golang/ds/tree/treeutils"
)

func main() {
	var root = &tree.Node{Data: 1}

	root.Left = &tree.Node{Data: 2}
	root.Right = &tree.Node{Data: 3}
	root.Right.Left = &tree.Node{Data: 4}
	root.Right.Right = &tree.Node{Data: 5}

	root.Right.Left.Left = &tree.Node{Data: 6}
	root.Right.Left.Right = &tree.Node{Data: 7}
	root.Right.Right.Left = &tree.Node{Data: 8}
	root.Right.Right.Right = &tree.Node{Data: 9}

	root.Left.Left = &tree.Node{Data: 10}
	root.Left.Right = &tree.Node{Data: 11}

	root.Left.Left.Right = &tree.Node{Data: 12}
	root.Left.Left.Left = &tree.Node{Data: 13}

	root.Left.Right.Left = &tree.Node{Data: 14}
	root.Left.Right.Right = &tree.Node{Data: 15}

	// fmt.Println("--PreOrder--")
	// treeutils.PreOrder(root)
	// fmt.Println("\n--InOrder--")
	// treeutils.InOrder(root)
	// fmt.Println("\n--PostOrder--")
	// treeutils.PostOrder(root)
	// fmt.Println("\n--level order--")
	// treeutils.LevelOrder(root)

	var count = treeutils.NodesCountIterative(root)
	fmt.Printf("Total node: %d\n", count)

	treeutils.TopViewOfTree(root)
}
