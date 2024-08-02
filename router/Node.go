package router

import "sync"

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
