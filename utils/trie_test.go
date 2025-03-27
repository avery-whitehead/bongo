package utils

import (
	"slices"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewTrie("")
	value := trie.Value
	children := trie.Children

	if value != "" {
		t.Errorf("root node value empty string, was %s", value)
	}

	if children == nil || len(children) != 0 {
		t.Errorf("root node children should be empty, was %+v", children)
	}
}

func TestInsert(t *testing.T) {
	trie := NewTrie("")

	trie.Insert("abc")
	// Indexed by rune value
	a := trie.Children[97]
	b := trie.Children[97].Children[98]
	c := trie.Children[97].Children[98].Children[99]

	if a == nil || b == nil || c == nil || !c.IsEnd {
		t.Errorf("a: %+v, b: %+v, c: %+v", a, b, c)
	}
}

func TestFindWords(t *testing.T) {
	trie := NewTrie("")
	test_words := []string{ "apple", "tries", "pairs", "paris", "steps" }
	test_letters := []rune {'a', 'e', 'i', 'p', 'r', 's', 't'}
	expected_words := []string { "pairs", "paris", "steps", "tries" }

	for _, word := range test_words {
		trie.Insert(word)
	}

	words := trie.FindWords(test_letters)
	// Sort response alphabetically because the order needs to match but we
	// don't care what the order is
	slices.Sort(words)

	if !slices.Equal(expected_words, words) {
		t.Errorf("expected: %+v found: %+v", expected_words, words)
	}
}