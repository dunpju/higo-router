package router

import "sync"

type RoutesMap struct {
	sort *Sort[string]
	list sync.Map
}

func newRouteMap() *RoutesMap {
	return &RoutesMap{sort: newSort[string](), list: sync.Map{}}
}

func (this *RoutesMap) Put(key string, routes *Routes) {
	this.sort.Append(key)
	this.list.Store(key, routes)
}

func (this *RoutesMap) Get(key string) (*Routes, bool) {
	if value, ok := this.list.Load(key); ok {
		return value.(*Routes), ok
	}
	return nil, false
}

func (this *RoutesMap) Range(fn func(key string, value *Routes) bool) {
	this.sort.Range(func(index int, key string) bool {
		value, _ := this.Get(key)
		return fn(key, value)
	})
}

func (this *RoutesMap) Len() int {
	return this.sort.Len()
}

type Serve struct {
	sort []string
	list *RoutesMap
	lock *sync.Mutex
}

func NewServe(name string) *Serve {
	serve := &Serve{sort: make([]string, 0), list: newRouteMap(), lock: new(sync.Mutex)}
	serve.Append(NewRoutes(name))
	return serve
}

// AddServe 添加 serve
func AddServe(name string) *Routes {
	currentServe = name
	if routes, ok := serve.list.Get(name); ok {
		return routes
	}
	routes := NewRoutes(name)
	serve.Append(routes)
	return routes
}

func (this *Serve) Sort() []string {
	return this.sort
}

func (this *Serve) List() *RoutesMap {
	return this.list
}

// Append 追加 serve
func (this *Serve) Append(routes *Routes) *Serve {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.sort = append(this.sort, routes.serve)
	this.list.Put(routes.serve, routes)
	return this
}

// Exist 是否存在
func (this *Serve) Exist(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.list.Get(name)
	return ok
}

func (this *Serve) Routes(name string) *Routes {
	this.lock.Lock()
	defer this.lock.Unlock()
	routes, ok := this.list.Get(name)
	if !ok {
		panic("Serve non-existent")
	}
	return routes
}

// AddRoute 添加 route
func (this *Serve) AddRoute(name string, route *Route) *Serve {
	this.lock.Lock()
	defer this.lock.Unlock()
	routes, ok := this.list.Get(name)
	if !ok {
		panic("Serve non-existent")
	} else {
		routes.Append(route)
	}
	return this
}

// ForEach 遍历 list
func (this *Serve) ForEach(callable StringCallable) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, index := range this.sort {
		value, _ := this.list.Get(index)
		callable(index, value)
	}
}

// GetServes 获取 serves
func GetServes() *Serve {
	return serve
}
