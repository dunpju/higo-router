package router

import (
	"reflect"
	"runtime"
	"strings"
)

var (
	currentServe           string
	currentGroupPrefix     string
	currentGroupMiddleware []interface{}
)

func AddRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(httpMethod, relativePath, handler, attributes...)
}

func AddGroup(prefix string, callable interface{}, attributes ...*RouteAttribute) {
	addGroup(prefix, callable, attributes...)
}

func Get(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(GET, relativePath, handler, attributes...)
}

func Post(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(POST, relativePath, handler, attributes...)
}

func Put(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(PUT, relativePath, handler, attributes...)
}

func Delete(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(DELETE, relativePath, handler, attributes...)
}

func Patch(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(PATCH, relativePath, handler, attributes...)
}

func Head(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(HEAD, relativePath, handler, attributes...)
}

func addRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	route := NewRoute()
	route.serve = currentServe
	route.method = strings.ToUpper(httpMethod)
	route.groupPrefix = currentGroupPrefix
	route.relativePath = relativePath
	route.handle = handler
	route.groupMiddle = currentGroupMiddleware
	for _, attribute := range attributes {
		if attribute.Name == ROUTE_FLAG {
			route.flag = attribute.Value.(string)
		} else if attribute.Name == ROUTE_FRONTPATH {
			route.frontPath = attribute.Value.(string)
		} else if attribute.Name == ROUTE_DESC {
			route.desc = attribute.Value.(string)
		} else if attribute.Name == ROUTE_IS_STATIC {
			route.isStatic = attribute.Value.(bool)
		} else if attribute.Name == ROUTE_MIDDLEWARE {
			route.middleware = append(route.middleware, attribute.Value)
		} else if attribute.Name == ROUTE_SERVE {
			route.serve = attribute.Value.(string)
		}
	}

	if "" == route.flag {
		if handle, ok := route.handle.(string); ok {
			route.flag = handle
		} else if _, ok := route.handle.(int); ok {
			panic("handle Can't be int")
		} else if _, ok := route.handle.(int64); ok {
			panic("handle Can't be int64")
		} else {
			route.flag = runtime.FuncForPC(reflect.ValueOf(route.handle).Pointer()).Name()
		}
	}

	if "" == route.serve {
		route.serve = DefaultServe
	}

	route.fullPath = route.groupPrefix + route.relativePath

	CollectRoute(route)
}

func addGroup(prefix string, callable interface{}, attributes ...*RouteAttribute) {
	previousGroupPrefix := currentGroupPrefix
	previousGroupMiddle := currentGroupMiddleware

	currentGroupPrefix = previousGroupPrefix + prefix
	currentGroupMiddleware = append(currentGroupMiddleware, RouteAttributes(attributes).Find(ROUTE_GROUP_MIDDLE))

	if fun, ok := callable.(func()); ok {
		fun() // 执行
	}

	currentGroupPrefix = previousGroupPrefix
	currentGroupMiddleware = previousGroupMiddle
}
