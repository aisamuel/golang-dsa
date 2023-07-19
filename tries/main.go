package main

import "fmt"

// Node
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

// InitTrue will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, word := range w {
		// fmt.Println(word)
		charIndex := word - 'a'
		// fmt.Println(charIndex)
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, word := range w {
		charIndex := word - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()
	myTrie.Insert("samuel")
	myTrie.Insert("peter")
	fmt.Println(myTrie.Search("peter"))
	fmt.Println(myTrie.Search("joseph"))

}
