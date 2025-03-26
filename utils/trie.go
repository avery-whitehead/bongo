package utils

// Represents a trie/prefix tree data structure, each struct instance is a
// node in the trie
type TrieNode struct {
	// Pointer to parent node
	Parent *TrieNode
	// Map of child nodes keyed by their character, nil if node does not exist
	Children map[rune]*TrieNode
	// Value of this node
	Value rune
	// True if this node marks the end of a word
	IsEnd bool
}

// Creates a new node
func NewTrie(value rune) *TrieNode {
	return &TrieNode{ Value: value, Children: make(map[rune]*TrieNode) }
}

// Inserts a new word into the trie
func (node *TrieNode) Insert(word string) {
	for _, char := range word {
		child := node.Children[char]
		// If this character does not exist as a child of this node, create
		// the node and add it to this node's children, keyed by the character
		if child == nil {
			child = NewTrie(char)
			node.Children[char] = child
		}
		// Set the parent of the child node to this node
		child.Parent = node
		// Set the current node to the child node, moving down the trie
		node = child
	}
	// Once each character has been added, mark the current node as the
	// end of the word
	node.IsEnd = true
}

// Traverses from a node back through its parents and builds the
// string represented by that path
func (node *TrieNode) GetWord() string {
	// Terminate once we've reached the root node
	if node.Parent == nil {
		return ""
	}

	return node.Parent.GetWord() + string(node.Value)
}

// Given a set of letters, find all the words in the trie that can be
// made from those 
func (node *TrieNode) FindWords(letters []rune) []string {
	results := make([]string, 0)

	if node.IsEnd {
		results = append(results, node.GetWord())
	}

	for _, letter := range letters {
		if node.Children[letter] != nil {
			results = append(results, node.Children[letter].FindWords(letters)...)
		}
	}

	return results
}