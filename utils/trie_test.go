package utils

import (
	"slices"
	"testing"
)

func TestInsert(t *testing.T) {
	trie := NewTrie(0)

	trie.Insert("abc")
	a := trie.Children[97]
	b := trie.Children[97].Children[98]
	c := trie.Children[97].Children[98].Children[99]

	if a == nil || b == nil || c == nil || !c.IsEnd {
		t.Errorf("a: %+v, b: %+v, c: %+v", a, b, c)
	}
}

func TestFindWords(t *testing.T) {
	trie := NewTrie(0)
	test_words := []string{ "apple", "tries", "pairs", "paris", "steps", "muddy" }
	test_letters := []rune {'a', 'e', 'i', 'p', 'r', 's', 't'}
	expected_words := []string { "pairs", "paris", "steps", "tries" }

	for _, word := range test_words {
		trie.Insert(word)
	}

	words := trie.FindWords(test_letters)
	slices.Sort(words)

	if !slices.Equal(expected_words, words) {
		t.Errorf("expected: %+v found: %+v", expected_words, words)
	}
}