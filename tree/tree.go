package tree

// Node represent Tree's node.
type Node struct {
	data  int
	left  *Node
	right *Node
}

type Binary struct {
	root *Node
}

// Insert add new value to binary tree.
func (b *Binary) Insert(value int) *Node {
	return insert(b.root, value)
}

// Search search value in binary tree.
func (b *Binary) Search(value int) bool {
	return search(b.root, value)
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

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{
			data: value,
		}
	}

	if value > root.data {
		root.right = insert(root.right, value)
	} else {
		root.left = insert(root.left, value)
	}

	return root
}

func NewBinary() *Binary {
	return &Binary{
		root: &Node{},
	}
}
