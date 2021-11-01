package router

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

var (
	currentServe           string //当前服务
	currentGroupPrefix     string //当前组前缀
	currentGroupIsAuth     bool   //当前组是否鉴权(默认:false)
	currentGroupMiddleware []interface{}
)

func AddRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(httpMethod, relativePath, handler, attributes...)
}

func AddGroup(prefix string, callable func(), attributes ...*RouteAttribute) {
	addGroup(prefix, callable, attributes...)
}

func Ws(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(WEBSOCKET, relativePath, handler, attributes...)
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

func addRoute(method string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	route := newRoute()
	route.serve = currentServe
	route.method = strings.ToUpper(method)
	route.groupPrefix = currentGroupPrefix
	route.relativePath = relativePath
	route.handle = handler
	route.groupMiddle = currentGroupMiddleware
	for _, attribute := range attributes {
		if attribute.Name == RouteFlag {
			route.flag = attribute.Value.(string)
		} else if attribute.Name == RouteFrontpath {
			route.frontPath = attribute.Value.(string)
		} else if attribute.Name == RouteDesc {
			route.desc = attribute.Value.(string)
		} else if attribute.Name == RouteIsStatic {
			route.isStatic = attribute.Value.(bool)
		} else if attribute.Name == RouteIsAuth {
			route.isAuth = attribute.Value.(bool)
		} else if attribute.Name == RouteMiddleware {
			route.middleware = append(route.middleware, attribute.Value)
		} else if attribute.Name == RouteServe {
			route.serve = attribute.Value.(string)
		} else if attribute.Name == RouteHeader {
			route.header = attribute.Value.(http.Header)
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

	route.absolutePath = route.groupPrefix + route.relativePath

	CollectRoute(route)
}

func addGroup(prefix string, callable func(), attributes ...*RouteAttribute) {
	previousGroupPrefix := currentGroupPrefix
	previousGroupIsAuth := currentGroupIsAuth
	previousGroupMiddle := currentGroupMiddleware

	currentGroupPrefix = previousGroupPrefix + prefix
	currentGroupMiddleware = append(currentGroupMiddleware, RouteAttributes(attributes).Find(RouteGroupMiddle))

	callable() // 执行

	currentGroupPrefix = previousGroupPrefix
	currentGroupIsAuth = previousGroupIsAuth
	currentGroupMiddleware = previousGroupMiddle
}
