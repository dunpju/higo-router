package router

import (
	"strings"
)

type RoutesCallable func(index int, route *Route)

type Routes struct {
	unique *UniqueString
	list   []*Route
}

func NewRoutes() *Routes {
	return &Routes{unique: NewUniqueString(), list: make([]*Route, 0)}
}

func (this *Routes) ForEach(callable RoutesCallable) {
	for key, value := range this.list {
		callable(key, value)
	}
}

// 追加 route
func (this *Routes) Append(route *Route) *Routes {
	this.unique.Append(route.unique)
	this.list = append(this.list, route)
	return this
}

// 收集 route
func CollectRoute(route *Route) {
	route.method = strings.ToUpper(route.method)
	if ! onlySupportMethods.Exist(route.method) {
		panic("route " + route.method + " error, only support:" + onlySupportMethods.String())
	}

	// 生成唯一标识
	route.UniMd5()

	if serve.Routes(route.serve).Unique().Exist(route.unique) {
		panic("route " + route.method + ":" + route.fullPath + " already exist")
	}

	serve.Routes(route.serve).Unique().Append(route.unique)
	serve.AddRoute(route.serve, route)
}

func (this *Routes) Unique() *UniqueString {
	return this.unique
}

func (this *Routes) List() []*Route {
	return this.list
}

// 获取路由集
func GetRoutes(name string) *Routes {
	return serve.Routes(name)
}

func (this *Routes) AddRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(httpMethod, relativePath, handler, attributes...)
	return this
}

func (this *Routes) AddGroup(prefix string, callable interface{}, attributes ...*RouteAttribute) *Routes {
	addGroup(prefix, callable, attributes...)
	return this
}

func (this *Routes) Get(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(GET, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Post(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(POST, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Put(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(PUT, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Delete(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(DELETE, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Patch(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(PATCH, relativePath, handler, attributes...)
	return this
}

func (this *Routes) Head(relativePath string, handler interface{}, attributes ...*RouteAttribute) *Routes {
	addRoute(HEAD, relativePath, handler, attributes...)
	return this
}
