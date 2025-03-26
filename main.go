package main

import (
	"fmt"
	"log"
	"os"

	"github.com/avery-whitehead/bongo/utils"
)

func main() {
	trie := utils.NewTrie(0, nil)

	_, err := os.ReadFile("./five_letter_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, word := range []string{"pairs", "paris", "apple", "tapes", "stair", "brown"} {
		trie.Insert(word)
	}
	//for word := range strings.SplitSeq(string(file), "\n") {
	//	trie.Insert(word)
	//}

	words := trie.FindWords([]rune{'a', 'e', 'i', 'r', 'p', 't', 's'})
	
	for _, word := range words {
		fmt.Println(word)
	}
}

