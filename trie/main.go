package main

import (
	"encoding/json"
	"fmt"
)

type trieNode struct {
	IsLeafe  bool               `json:"is_leafe"`
	Children map[rune]*trieNode `json:"children"`
}

func NewTrieNode() *trieNode {
	return &trieNode{
		IsLeafe:  false,
		Children: make(map[rune]*trieNode),
	}
}

func (t *trieNode) insert(key string) {
	curr := t
	for _, c := range key {
		if _, ok := curr.Children[c-'a']; !ok {
			curr.Children[c-'a'] = NewTrieNode()
		}

		curr = curr.Children[c-'a']
	}

	curr.IsLeafe = true
}

func (t *trieNode) search(key string) bool {
	curr := t
	for _, c := range key {
		if _, ok := curr.Children[c-'a']; !ok {
			return false
		}

		curr = curr.Children[c-'a']
	}

	if curr.IsLeafe == false {
		return false
	}

	return true
}

func (t *trieNode) print() {
	b, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

func main() {
	t := NewTrieNode()
	t.insert("abc")
	t.insert("abd")
	t.insert("a")

	found := t.search("abc")

	t.print()
	fmt.Println(found)
}
