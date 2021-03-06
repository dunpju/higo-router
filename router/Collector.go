package router

var (
	currentGroupPrefix     string
	currentGroupMiddleware interface{}
)

func AddRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute(httpMethod, relativePath, handler, attributes...)
}

func AddGroup(prefix string, callable interface{}, attributes ...*RouteAttribute) {
	addGroup(prefix, callable, attributes...)
}

func Get(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute("GET", relativePath, handler, attributes...)
}

func Post(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute("POST", relativePath, handler, attributes...)
}

func Put(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute("PUT", relativePath, handler, attributes...)
}

func Delete(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute("DELETE", relativePath, handler, attributes...)
}

func Patch(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute("PATCH", relativePath, handler, attributes...)
}

func Head(relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	addRoute("HEAD", relativePath, handler, attributes...)
}

func addRoute(httpMethod string, relativePath string, handler interface{}, attributes ...*RouteAttribute) {
	route := &Route{}
	route.groupPrefix = currentGroupPrefix
	route.method = httpMethod
	route.relativePath = currentGroupPrefix + relativePath
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
		} else if attribute.Name == ROUTE_MIDDLEWARE {
			route.middleware = attribute.Value
		}
	}
	AppendRoutes(route)
}

func addGroup(prefix string, callable interface{}, attributes ...*RouteAttribute) {
	previousGroupPrefix := currentGroupPrefix
	currentGroupMiddle := currentGroupMiddleware

	currentGroupPrefix = previousGroupPrefix + prefix
	currentGroupMiddleware = RouteAttributes(attributes).Find(ROUTE_MIDDLEWARE)

	if fun, ok := callable.(func()); ok {
		fun() // 执行
	}

	currentGroupPrefix = previousGroupPrefix
	currentGroupMiddleware = currentGroupMiddle
}
