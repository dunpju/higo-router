package router

import (
	"fmt"
	"strings"
)

type Node struct {
	isEnd    bool
	Children map[string]*Node
}

func NewNode() *Node {
	return &Node{Children: make(map[string]*Node)}
}

func (this *Node) Each(fu func(n *Node)) {
	fu(this)
	if len(this.Children) > 0 {
		for _, node := range this.Children {
			node.Each(fu)
		}
	}
}

type Trie struct {
	node *Node
}

func NewTrie() *Trie {
	return &Trie{node: NewNode()}
}

func (this *Trie) each(str string, fu func(s string)) {
	strs := strings.Split(str, "/")
	for _, s := range strs {
		fu(s)
	}
}

func (this *Trie) Each(fu func(n *Node)) {
	this.node.Each(fu)
}

func (this *Trie) Insert(str string) *Trie {
	current := this.node
	this.each(str, func(s string) {
		if _, ok := current.Children[s]; !ok {
			current.Children[s] = NewNode()
		}
		current = current.Children[s]
	})
	current.isEnd = true
	return this
}

func (this *Trie) Has(str string) bool {
	current := this.node
	for _, item := range ([]rune)(str) {
		if _, ok := current.Children[string(item)]; !ok {
			return false
		}
		current = current.Children[string(item)]
	}
	return current.isEnd
}

func (this *Trie) Search(str string) (*Node, error) {
	current := this.node
	for _, item := range ([]rune)(str) {
		if _, ok := current.Children[string(item)]; !ok {
			return nil, fmt.Errorf("not found")
		}
		current = current.Children[string(item)]
	}
	if current.isEnd {
		return current, nil
	}
	return nil, fmt.Errorf("not found")
}
