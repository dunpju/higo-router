package router

import "sync"

const (
	ROUTE_METHOD        = "method"
	ROUTE_RELATIVE_PATH = "relativePath"
	ROUTE_HANDLE        = "handle"
	ROUTE_FLAG          = "flag"
	ROUTE_FRONTPATH     = "frontPath"
	ROUTE_IS_STATIC     = "isStatic"
	ROUTE_DESC          = "desc"
	ROUTE_MIDDLEWARE    = "middleware"
)

type Route struct {
	method       string      // 请求方法 GET/POST/DELETE/PATCH/OPTIONS/HEAD
	relativePath string      // 后端 api relativePath
	handle       interface{} // 后端控制器函数
	flag         string      // 后端控制器函数标记
	frontPath    string      // 前端 path(前端菜单路由)
	isStatic     bool        // 是否静态文件
	desc         string      // 描述
	middleware   interface{} // 中间件
}

type Routes []*Route

var routes Routes
var routesOnce sync.Once

func AppendRoutes(route *Route) {
	routesOnce.Do(func() {
		routes = make(Routes, 0)
	})
	routes = append(routes, route)
}

// 获取路由集
func GetRoutes() Routes {
	return routes
}

type RouteAttributes []*RouteAttribute

func (this RouteAttributes) Find(name string) interface{} {
	for _, p := range this {
		if p.Name == name {
			return p.Value
		}
	}
	return nil
}

type RouteAttribute struct {
	Name  string
	Value interface{}
}

func NewRouteAttribute(name string, value interface{}) *RouteAttribute {
	return &RouteAttribute{Name: name, Value: value}
}

func Flag(value string) *RouteAttribute {
	return NewRouteAttribute(ROUTE_FLAG, value)
}

func FrontPath(value string) *RouteAttribute {
	return NewRouteAttribute(ROUTE_FRONTPATH, value)
}

func IsStatic(value bool) *RouteAttribute {
	return NewRouteAttribute(ROUTE_IS_STATIC, value)
}

func Desc(value string) *RouteAttribute {
	return NewRouteAttribute(ROUTE_DESC, value)
}

func Middleware(value string) *RouteAttribute {
	return NewRouteAttribute(ROUTE_MIDDLEWARE, value)
}
