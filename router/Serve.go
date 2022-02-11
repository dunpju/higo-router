package router

import "sync"

// serve
type Serve struct {
	sort []string
	list map[string]*Routes
	lock *sync.Mutex
}

func NewServe(name string) *Serve {
	serve := &Serve{sort: make([]string, 0), list: make(map[string]*Routes), lock: new(sync.Mutex)}
	serve.Append(NewRoutes(name))
	return serve
}

// 添加 serve
func AddServe(name string) *Routes {
	currentServe = name
	if routes, ok := serve.list[name]; ok {
		return routes
	}
	routes := NewRoutes(name)
	serve.Append(routes)
	return routes
}

func (this *Serve) Sort() []string {
	return this.sort
}

func (this *Serve) List() map[string]*Routes {
	return this.list
}

// 追加 serve
func (this *Serve) Append(routes *Routes) *Serve {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.sort = append(this.sort, routes.serve)
	this.list[routes.serve] = routes
	return this
}

// 是否存在
func (this *Serve) Exist(name string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, ok := this.list[name]
	return ok
}

func (this *Serve) Routes(name string) *Routes {
	this.lock.Lock()
	defer this.lock.Unlock()
	routes, ok := this.list[name]
	if ! ok {
		panic("Serve non-existent")
	}
	return routes
}

// 添加 route
func (this *Serve) AddRoute(name string, route *Route) *Serve {
	this.lock.Lock()
	defer this.lock.Unlock()
	routes, ok := this.list[name]
	if ! ok {
		panic("Serve non-existent")
	} else {
		routes.Append(route)
	}
	return this
}

// 遍历 list
func (this *Serve) ForEach(callable StringCallable) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, index := range this.sort {
		callable(index, this.list[index])
	}
}

// 获取 serves
func GetServes() *Serve {
	return serve
}
