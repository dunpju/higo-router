package router

import (
	"fmt"
	"strings"
)

type Node struct {
	isEnd    bool
	Route    *Route
	Children map[string]*Node
}

func (this *Node) IsEnd() bool {
	return this.isEnd
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

func (this *Trie) each(str string, fu func(s string) bool) {
	strs := strings.Split(str, "/")
	for _, s := range strs {
		if !fu(s) {
			break
		}
	}
}

func (this *Trie) Each(fu func(n *Node)) {
	this.node.Each(fu)
}

func (this *Trie) insert(route *Route) *Trie {
	str := route.absolutePath
	current := this.node
	this.each(str, func(s string) bool {
		if _, ok := current.Children[s]; !ok {
			current.Children[s] = NewNode()
		}
		current = current.Children[s]
		return true
	})
	current.Route = route
	current.isEnd = true
	return this
}

func (this *Trie) Insert(str string) *Trie {
	current := this.node
	this.each(str, func(s string) bool {
		if _, ok := current.Children[s]; !ok {
			current.Children[s] = NewNode()
		}
		current = current.Children[s]
		return true
	})
	current.isEnd = true
	return this
}

func (this *Trie) Has(str string) bool {
	current := this.node
	this.each(str, func(s string) bool {
		if _, ok := current.Children[s]; !ok {
			return false
		}
		current = current.Children[s]
		return true
	})
	return current.isEnd
}

func (this *Trie) Search(str string) (*Node, error) {
	current := this.node
	this.each(str, func(s string) bool {
		if _, ok := current.Children[s]; !ok {
			current = nil
			return false
		}
		current = current.Children[s]
		return true
	})
	if current == nil {
		return nil, fmt.Errorf("not found")
	}
	if current.isEnd {
		return current, nil
	}
	return nil, fmt.Errorf("not found")
}
