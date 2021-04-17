package router

// serve
type Serve struct {
	sort []string
	list map[string]*Routes
}

func NewServe(name string) *Serve {
	ser := &Serve{sort: make([]string, 0), list: make(map[string]*Routes)}
	ser.Append(name, NewRoutes())
	return ser
}

// 添加 serve
func AddServe(name string)  {
	serve.Append(name, NewRoutes())
}

func (this *Serve) Sort() []string {
	return this.sort
}

func (this *Serve) List() map[string]*Routes {
	return this.list
}

// 追加 serve
func (this *Serve) Append(name string, routes *Routes) *Serve {
	this.sort = append(this.sort, name)
	this.list[name] = routes
	return this
}

// 是否存在
func (this *Serve) Exist(name string) bool {
	_, ok := this.list[name]
	return ok
}

func (this *Serve) Routes(name string) *Routes {
	routes, ok := this.list[name]
	if ! ok {
		panic("Serve non-existent")
	}
	return routes
}

// 添加 route
func (this *Serve) AddRoute(name string, route *Route) *Serve {
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
	for _, index := range this.sort {
		callable(index, this.list[index])
	}
}

// 获取 serves
func GetServes() *Serve {
	return serve
}
