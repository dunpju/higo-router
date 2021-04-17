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
