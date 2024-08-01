package router

import (
	"fmt"
	"strings"
	"sync"
)

type Sort[T int | string] struct {
	sort []T
}

func newSort[T int | string]() *Sort[T] {
	return &Sort[T]{sort: make([]T, 0)}
}

func (this *Sort[T]) Exist(value T) bool {
	for _, v := range this.sort {
		if v == value {
			return true
		}
	}
	return false
}

func (this *Sort[T]) Append(value T) {
	if !this.Exist(value) {
		this.sort = append(this.sort, value)
	}
}

func (this *Sort[T]) Range(fn func(index int, value T) bool) {
	for i, v := range this.sort {
		if !fn(i, v) {
			break
		}
	}
}

func (this *Sort[T]) Len() int {
	return len(this.sort)
}

type NodeMap struct {
	sort *Sort[string]
	list sync.Map
}

func newNodeMap() *NodeMap {
	return &NodeMap{sort: newSort[string](), list: sync.Map{}}
}

func (this *NodeMap) Put(key string, node *Node) {
	this.sort.Append(key)
	this.list.Store(key, node)
}

func (this *NodeMap) Get(key string) (*Node, bool) {
	if node, ok := this.list.Load(key); ok {
		return node.(*Node), ok
	}
	return nil, false
}

func (this *NodeMap) Range(fn func(key string, node *Node) bool) {
	this.sort.Range(func(index int, value string) bool {
		node, _ := this.Get(value)
		return fn(value, node)
	})
}

func (this *NodeMap) Len() int {
	return this.sort.Len()
}

type ParamMap struct {
	sort *Sort[int]
	list sync.Map
}

func newParamMap() *ParamMap {
	return &ParamMap{sort: newSort[int](), list: sync.Map{}}
}

func (this *ParamMap) Put(key int, value []string) {
	this.sort.Append(key)
	this.list.Store(key, value)
}

func (this *ParamMap) Append(key int, value string) {
	values, _ := this.Get(key)
	values = append(values, value)
	this.list.Store(key, values)
}

func (this *ParamMap) Get(key int) ([]string, bool) {
	if value, ok := this.list.Load(key); ok {
		return value.([]string), ok
	}
	return nil, false
}

func (this *ParamMap) Range(fn func(key int, value []string) bool) {
	this.sort.Range(func(index int, key int) bool {
		value, _ := this.Get(key)
		return fn(key, value)
	})
}

func (this *ParamMap) Len() int {
	return this.sort.Len()
}

type Node struct {
	isEnd    bool
	hasParam bool
	suffix   string
	Param    *ParamMap
	Route    *Route
	Children *NodeMap
}

func (this *Node) HasParam() bool {
	return this.hasParam
}

func (this *Node) IsEnd() bool {
	return this.isEnd
}

func NewNode() *Node {
	return &Node{Param: newParamMap(), Children: newNodeMap()}
}

func (this *Node) Each(fu func(n *Node)) {
	fu(this)
	if this.Children.Len() > 0 {
		this.Children.Range(func(key string, node *Node) bool {
			node.Each(fu)
			return true
		})
	}
}

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
		return nil, fmt.Errorf("not found")
	}
	if current.isEnd {
		return current, nil
	}
	return nil, fmt.Errorf("not found")
}
