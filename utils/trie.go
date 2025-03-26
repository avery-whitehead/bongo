package utils

type Trie struct {
	Parent *Trie
	Children map[rune]*Trie
	Value rune
	IsEnd bool
}

func NewTrie() *Trie {
	return &Trie{ Parent: nil, Children: make(map[rune]*Trie), IsEnd: false }
}

func (trie *Trie) Insert(word string) {
	node := trie
	for _, char := range word {
		child := node.Children[char]
		if child == nil {
			if node.Children == nil {
				node.Children = make(map[rune]*Trie)
			}
			child = new(Trie)
			node.Children[char] = child
		}
		child.Parent = node
		node = child
		node.Value = char
	}
	node.IsEnd = true
}

func (trie *Trie) GetWord() string {
	if trie.Parent == nil {
		return string(trie.Value)
	}

	return trie.Parent.GetWord() + string(trie.Value)
}

func (trie *Trie) FindWords(letters []rune) []string {
	results := make([]string, 0)

	if trie.IsEnd {
		results = append(results, trie.GetWord())
	}

	for _, letter := range letters {
		if trie.Children[letter] != nil {
			results = append(results, trie.Children[letter].FindWords(letters)...)
		}
	}


	return results
}