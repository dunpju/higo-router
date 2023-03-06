package router

import "fmt"

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (this *Trie) Insert(str string) {
	current := this.root
	for _, item := range ([]rune)(str) {
		if _, ok := current.Children[string(item)]; !ok {
			current.Children[string(item)] = NewNode()
		}
		current = current.Children[string(item)]
	}
	current.isEnd = true
}

func (this *Trie) Has(str string) bool {
	current := this.root
	for _, item := range ([]rune)(str) {
		if _, ok := current.Children[string(item)]; !ok {
			return false
		}
		current = current.Children[string(item)]
	}
	return current.isEnd
}

func (this *Trie) Search(str string) (*Node, error) {
	current := this.root
	for _, item := range ([]rune)(str) {
		if node, ok := current.Children[string(item)]; !ok {
			return node, fmt.Errorf("not found")
		}
		current = current.Children[string(item)]
	}
	panic(fmt.Errorf("not found"))
}
