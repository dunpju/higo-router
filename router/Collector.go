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
	currentGroupIsDataAuth bool   //当前组是否数据鉴权(默认:false)
	currentGroupMiddleware []interface{}
	globalGroupPrefix      string //全局组前缀
	globalApiGroupPrefix   string //全局API组前缀
)

func GlobalGroupIsAuth(b bool) {
	currentGroupIsAuth = b
}

func GlobalGroupPrefix(prefix string) {
	globalGroupPrefix = prefix
}

func GlobalApiGroupPrefix(prefix string) {
	globalApiGroupPrefix = prefix
}

func GlobalGroupIsDataAuth(b bool) {
	currentGroupIsDataAuth = b
}

func AddRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(httpMethod, relativePath, handler, attributes...)
}

func AddGroup(prefix string, callable func(), attributes ...*RouteAttribute) {
	addGroup(prefix, callable, attributes...)
}

func Ws(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(WEBSOCKET, relativePath, handler, append(attributes, IsWs(true))...)
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
	route.groupPrefix = globalGroupPrefix + globalApiGroupPrefix + currentGroupPrefix
	route.isAuth = currentGroupIsAuth
	route.isDataAuth = currentGroupIsDataAuth
	route.relativePath = relativePath
	route.handle = handler
	route.groupMiddle = append(route.groupMiddle, currentGroupMiddleware...)
	for _, attribute := range attributes {
		if attribute.Name == RouteFlag {
			route.flag = attribute.Value[0].(string)
		} else if attribute.Name == RouteFrontpath {
			route.frontPath = attribute.Value[0].(string)
		} else if attribute.Name == RouteDesc {
			route.desc = attribute.Value[0].(string)
		} else if attribute.Name == RouteTitle {
			route.title = attribute.Value[0].(string)
		} else if attribute.Name == RouteServe {
			route.serve = attribute.Value[0].(string)
		} else if attribute.Name == RouteIsStatic {
			route.groupPrefix = globalGroupPrefix + currentGroupPrefix
			route.isStatic = attribute.Value[0].(bool)
		} else if attribute.Name == RouteCancelGlobalGroupPrefix {
			if attribute.Value[0].(bool) {
				route.groupPrefix = strings.Replace(route.groupPrefix, globalGroupPrefix, "", 1)
			}
		} else if attribute.Name == RouteCancelGlobalApiGroupPrefix {
			if attribute.Value[0].(bool) {
				route.groupPrefix = strings.Replace(route.groupPrefix, globalApiGroupPrefix, "", 1)
			}
		} else if attribute.Name == RouteIsAuth {
			route.isAuth = attribute.Value[0].(bool)
		} else if attribute.Name == RouteIsDataAuth {
			route.isDataAuth = attribute.Value[0].(bool)
		} else if attribute.Name == RouteIsWs {
			route.isWs = attribute.Value[0].(bool)
		} else if attribute.Name == RouteMiddleware {
			route.middleware = append(route.middleware, attribute.Value...)
		} else if attribute.Name == RouteGlobalMiddle {
			route.globalMiddle = append(route.globalMiddle, attribute.Value...)
		} else if attribute.Name == RouteHeader {
			route.header = attribute.Value[0].(http.Header)
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
	previousGroupIsDataAuth := currentGroupIsDataAuth
	previousGroupMiddle := currentGroupMiddleware

	currentGroupPrefix = previousGroupPrefix + prefix
	if isAuth := RouteAttributes(attributes).Find(RouteIsAuth); isAuth != nil {
		currentGroupIsAuth = isAuth.(bool)
	}
	if isDataAuth := RouteAttributes(attributes).Find(RouteIsDataAuth); isDataAuth != nil {
		currentGroupIsDataAuth = isDataAuth.(bool)
	}

	if nil != RouteAttributes(attributes).Find(RouteGroupMiddle) {
		currentGroupMiddleware = append(currentGroupMiddleware, RouteAttributes(attributes).Find(RouteGroupMiddle).([]interface{})...)
	}

	callable() // 执行

	currentGroupPrefix = previousGroupPrefix
	currentGroupIsAuth = previousGroupIsAuth
	currentGroupIsDataAuth = previousGroupIsDataAuth
	currentGroupMiddleware = previousGroupMiddle
}
