package trie

const (
	alphabetSize = 26
)

// Node represent trie's node.
type Node struct {
	childern [alphabetSize]*Node
	isEnd    bool
}

//  Trie represent Trie data structure.
type Trie struct {
	root *Node
}

// Insert add new word to trie.
func (t *Trie) Insert(word string) {
	var currentNode = t.root

	for _, value := range []rune(word) {
		var charIndex = value - 'a' // to get zero based index for char

		if currentNode.childern[charIndex] == nil {
			currentNode.childern[charIndex] = &Node{}
		}

		currentNode = currentNode.childern[charIndex]
	}
	currentNode.isEnd = true
}

//  Search search given word in trie and return true
//  if word is present in trie.
func (t *Trie) Search(word string) bool {
	var currentNode = t.root

	for _, value := range []rune(word) {
		var charIndex = value - 'a' // to get zero based index for char

		if currentNode.childern[charIndex] == nil {
			return false
		}
		// Update current node.
		currentNode = currentNode.childern[charIndex]
	}

	return currentNode.isEnd
}

// New create new trie object.
func New() *Trie {
	return &Trie{
		root: &Node{},
	}
}
