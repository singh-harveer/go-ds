package main

import (
	"fmt"
	"golang/ds/tree"
	"golang/ds/tree/bst"
	"golang/ds/tree/treeutils"
)

func main() {
	var root *tree.Node
	var binary = bst.NewBinary()
	for _, num := range []int{5, 4, 3, 2, 1, 6, 7, 8, 9} {
		root = binary.Insert(num)
	}
	// fmt.Println("--PreOrder--")
	// treeutils.PreOrder(root)
	// fmt.Println("\n--InOrder--")
	// treeutils.InOrder(root)
	// fmt.Println("\n--PostOrder--")
	// treeutils.PostOrder(root)
	// fmt.Println("\n--level order--")
	// treeutils.LevelOrder(root)
	// var count = treeutils.CountNodeRecursive(root)
	// fmt.Printf("Total node: %d\n", count)

	var count = treeutils.NodesCountIterative(root)
	fmt.Printf("Total node: %d\n", count)
}
