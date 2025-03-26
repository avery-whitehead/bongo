package utils

type Trie struct {
	Parent *Trie
	Value rune
	Children map[rune]*Trie
	IsEnd bool
}

func NewTrie(value rune, parent *Trie) *Trie {
	return &Trie { Value: value, Parent: parent, Children: make(map[rune]*Trie) }
}

func (trie *Trie) Insert(word string) {
	for _, char := range word {
		if _, ok := trie.Children[char]; !ok {
			trie.Children[char] = NewTrie(char, trie)
		} else {
			trie = trie.Children[char]
		}
	}
	trie.IsEnd = true
}

func (trie *Trie) GetWord() string {
	if trie.Parent == nil {
		return string(trie.Value)
	}
	
	return trie.Parent.GetWord() + string(trie.Value)
}

func (trie *Trie) FindWords(letters []rune) []string {
	var results []string

	if trie.IsEnd {
		results = append(results, trie.GetWord())
	}

	for _, letter := range letters {
		if _, ok := trie.Children[letter]; ok {
			//letterRemoved := slices.Delete(letters, i, i + 1)
			//results = append(results, trie.Children[letter].FindWords(letterRemoved)...)
			results = append(results, trie.Children[letter].FindWords((letters))...)
		}
	}

	return results
}