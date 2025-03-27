package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/avery-whitehead/bongo/utils"
)

func main() {
	trie := utils.NewTrie("")

	file, err := os.ReadFile("./five_letter_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	for word := range strings.SplitSeq(string(file), "\n") {
		trie.Insert(word)
	}

	words := trie.FindWords([]rune{'a', 'e', 'i', 'r', 'p', 't', 's'})
	
	for i, word := range words {
		fmt.Printf("%d: %s\n", i, word)
	}
}

