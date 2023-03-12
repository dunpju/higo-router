package router

import (
	"strings"
	"sync"
)

type RoutesCallable func(route *Route)

type Routes struct {
	serve string
	trie  *Trie
	lock  *sync.Mutex
}

func (this *Routes) Trie() *Trie {
	return this.trie
}

func NewRoutes(name string) *Routes {
	return &Routes{serve: name, trie: NewTrie(), lock: new(sync.Mutex)}
}

func (this *Routes) ForEach(callable RoutesCallable) {
	this.trie.Each(func(n *Node) {
		if n.Route != nil {
			callable(n.Route)
		}
	})
}

func (this *Routes) Search(method, str string) (*Node, error) {
	return this.trie.Search(method, str)
}

func (this *Routes) Route(method, url string) (*Route, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if node, err := this.trie.Search(method, url); err == nil {
		return node.Route, nil
	} else {
		return nil, err
	}
}

func (this *Routes) Exist(method, url string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, err := this.trie.Search(method, url); err == nil {
		return true
	} else {
		return false
	}
}

// 追加 route
func (this *Routes) Append(route *Route) *Routes {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.trie.insert(route)
	return this
}

// 收集 route
func CollectRoute(route *Route) {
	route.method = strings.ToUpper(route.method)
	if !onlySupportMethods.Exist(route.method) {
		panic(route.serve + " route " + route.method + " error, only support:" + onlySupportMethods.String())
	}

	serve.AddRoute(route.serve, route)
}

func (this *Routes) Serve() string {
	return this.serve
}

// 获取路由集
func GetRoutes(name string) *Routes {
	return serve.Routes(name)
}

func (this *Routes) AddRoute(method string, relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.lock.Lock()
	defer this.lock.Unlock()
	if nil == RouteAttributes(attributes).Find(RouteServe) {
		attributes = RouteAttributes(attributes).Append(SetServe(this.serve))
	}
	addRoute(method, relativePath, handler, attributes...)
	return this
}

func (this *Routes) AddGroup(prefix string, callable func(), attributes ...*RouteAttribute) *Routes {
	if nil == RouteAttributes(attributes).Find(RouteServe) {
		attributes = RouteAttributes(attributes).Append(SetServe(this.serve))
	}
	addGroup(prefix, callable, attributes...)
	return this
}

func (this *Routes) Ws(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(WEBSOCKET, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Get(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(GET, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Post(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(POST, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Put(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(PUT, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Delete(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(DELETE, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Patch(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(PATCH, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Head(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	this.AddRoute(HEAD, relativePath, handler, attributes...)
	return this
}
