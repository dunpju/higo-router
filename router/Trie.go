package router

import (
	"fmt"
	"strings"
)

type Node struct {
	isEnd    bool
	hasParam bool
	suffix   string
	Param    map[int][]string
	Route    *Route
	Children map[string]*Node
}

func (this *Node) HasParam() bool {
	return this.hasParam
}

func (this *Node) IsEnd() bool {
	return this.isEnd
}

func NewNode() *Node {
	return &Node{Param: make(map[int][]string), Children: make(map[string]*Node)}
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
	node map[string]*Node
}

func NewTrie() *Trie {
	return &Trie{node: make(map[string]*Node)}
}

func (this *Trie) split(str string, fu func(s string) bool) {
	strs := strings.Split(str, "/")
	for _, s := range strs {
		if !fu(s) {
			break
		}
	}
}

func (this *Trie) Each(fu func(n *Node)) {
	for _, node := range this.node {
		node.Each(fu)
	}
}

func (this *Trie) insert(route *Route) *Trie {
	if !onlySupportMethods.Exist(route.method) {
		panic(route.serve + " route " + route.method + " error, only support:" + onlySupportMethods.String())
	}
	if this.Has(route.method, route.absolutePath) {
		panic(route.serve + " route " + route.method + ":" + route.absolutePath + " already exist")
	}
	current, ok := this.node[route.method]
	if !ok {
		current = NewNode()
		this.node[route.method] = current
	}

	str := route.absolutePath
	params := make([]string, 0)

	this.split(str, func(s string) bool {
		if s != "" {
			if string(s[0]) == ":" {
				params = append(params, s)
				return true
			}
		}
		if _, ok := current.Children[s]; !ok {
			n := NewNode()
			current.Children[s] = n
		}
		if len(params) > 0 {
			current.suffix = "/"
			if _, ok := current.Param[len(params)]; ok {
				current.Param[len(params)] = append(current.Param[len(params)], current.suffix+s)
			} else {
				current.Param[len(params)] = []string{current.suffix + s}
			}
			current.hasParam = true
		}
		current = current.Children[s]
		return true
	})
	if len(params) > 0 && !current.hasParam {
		current.suffix = "/"
		if _, ok := current.Param[len(params)]; ok {
			current.Param[len(params)] = append(current.Param[len(params)], current.suffix)
		} else {
			current.Param[len(params)] = []string{current.suffix}
		}
		current.hasParam = true
	}
	current.Route = route
	current.isEnd = true
	return this
}

func (this *Trie) Has(method, str string) bool {
	if _, ok := this.node[method]; !ok {
		return ok
	}
	if _, err := this.Search(method, str); err != nil {
		return false
	}
	return true
}

func (this *Trie) Search(method, str string) (*Node, error) {
	current := this.node[method]
	paramCounter := 0
	this.split(str, func(s string) bool {
		if _, ok := current.Children[s]; !ok {
			if current.hasParam {
				paramCounter++
			} else {
				current = nil
				return false
			}
		} else {
			if current.hasParam {
				if _, ok := current.Param[paramCounter]; !ok && paramCounter != 0 {
					current = nil
					return false
				}
			}
			current = current.Children[s]
		}
		return true
	})
	if current != nil && current.hasParam {
		if params, ok := current.Param[paramCounter]; !ok && paramCounter != 0 {
			current = nil
		} else {
			hasSuffix := false
			for _, p := range params {
				if current.suffix == p {
					hasSuffix = true
				}
			}
			if !hasSuffix {
				current = nil
			}
		}
	}
	if current == nil {
		return nil, fmt.Errorf("not found")
	}
	if current.isEnd {
		return current, nil
	}
	return nil, fmt.Errorf("not found")
}
