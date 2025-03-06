package router

import (
	"fmt"
	"strings"
)

type Trie struct {
	node *NodeMap
}

func NewTrie() *Trie {
	return &Trie{node: newNodeMap()}
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
	this.node.Range(func(key string, node *Node) bool {
		node.Each(fu)
		return true
	})
}

func (this *Trie) insert(route *Route) *Trie {
	if !onlySupportMethods.Exist(route.method) {
		panic(route.serve + " route " + route.method + " error, only support:" + onlySupportMethods.String())
	}
	if this.Has(route.method, route.absolutePath) {
		panic(route.serve + " route " + route.method + ":" + route.absolutePath + " already exist")
	}
	current, ok := this.node.Get(route.method)
	if !ok {
		current = NewNode()
		this.node.Put(route.method, current)
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
		if _, ok := current.Children.Get(s); !ok {
			n := NewNode()
			current.Children.Put(s, n)
		}
		if len(params) > 0 {
			current.suffix = "/"
			if _, ok := current.Param.Get(len(params)); ok {
				current.Param.Append(len(params), current.suffix+s)
			} else {
				current.Param.Put(len(params), []string{current.suffix + s})
			}
			current.hasParam = true
		}
		current, _ = current.Children.Get(s)
		return true
	})
	if len(params) > 0 && !current.hasParam {
		current.suffix = "/"
		if _, ok := current.Param.Get(len(params)); ok {
			current.Param.Append(len(params), current.suffix)
		} else {
			current.Param.Put(len(params), []string{current.suffix})
		}
		current.hasParam = true
	}
	current.Route = route
	current.isEnd = true
	return this
}

func (this *Trie) Has(method, str string) bool {
	if _, ok := this.node.Get(method); !ok {
		return ok
	}
	if _, err := this.Search(method, str); err != nil {
		return false
	}
	return true
}

func (this *Trie) Search(method, str string) (*Node, error) {
	current, _ := this.node.Get(method)
	if current == nil {
		return nil, fmt.Errorf(method + " not found")
	}
	paramCounter := 0
	this.split(str, func(s string) bool {
		if _, ok := current.Children.Get(s); !ok {
			if current.hasParam {
				paramCounter++
			} else {
				current = nil
				return false
			}
		} else {
			if current.hasParam {
				if _, ok := current.Param.Get(paramCounter); !ok && paramCounter != 0 {
					current = nil
					return false
				}
			}
			current, _ = current.Children.Get(s)
			if current == nil {
				return false
			}
		}
		return true
	})
	if current != nil && current.hasParam {
		if params, ok := current.Param.Get(paramCounter); !ok && paramCounter != 0 {
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
		return nil, fmt.Errorf("current not found")
	}
	if current.isEnd {
		return current, nil
	}
	return nil, fmt.Errorf("not found")
}
