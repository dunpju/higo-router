package router

import (
	"fmt"
)

var (
	currentGroupPrefix string
)

func AddRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	fmt.Println("AddRoute:" + currentGroupPrefix, relativePath, handler)
	route := &Route{}
	route.method = httpMethod
	route.relativePath = relativePath
	route.handle = handler
	for _, attribute := range attributes {
		if attribute.Name == ROUTE_FLAG {
			route.flag = attribute.Value.(string)
		} else if attribute.Name == ROUTE_FRONTPATH {
			route.frontPath = attribute.Value.(string)
		} else if attribute.Name == ROUTE_DESC {
			route.desc = attribute.Value.(string)
		} else if attribute.Name == ROUTE_IS_STATIC {
			route.isStatic = attribute.Value.(bool)
		}
	}
	AppendRoutes(route)
}

func AddGroup(prefix string, callable interface{}) {
	previousGroupPrefix := currentGroupPrefix
	currentGroupPrefix = previousGroupPrefix + prefix
	if _, ok := callable.(func()); ok {
		callable.(func())() // 执行回调
	}
	currentGroupPrefix = previousGroupPrefix
}
