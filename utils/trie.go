package utils

// Represents a trie/prefix tree data structure, each struct instance is a
// node in the trie
type TrieNode struct {
	// Map of child nodes keyed by their character, nil if node does not exist
	Children map[rune]*TrieNode
	// Value of this node
	Value string
	// True if this node marks the end of a word
	IsEnd bool
}

// Creates a new node
func NewTrie(value string) *TrieNode {
	return &TrieNode{ Value: value, Children: make(map[rune]*TrieNode) }
}

// Inserts a new word into the trie
func (node *TrieNode) Insert(word string) {
	for _, char := range word {
		child := node.Children[char]
		if child == nil {
			child = NewTrie(node.Value + string(char))
			node.Children[char] = child
		}
		node = child
	}
	node.IsEnd = true
}

// Given a set of letters, find all the words in the trie that can be
// made from those 
func (node *TrieNode) FindWords(letters []rune) []string {
	results := make([]string, 0)

	if node.IsEnd {
		results = append(results, node.Value)
	}

	for _, letter := range letters {
		if node.Children[letter] != nil {
			results = append(results, node.Children[letter].FindWords(letters)...)
		}
	}

	return results
}