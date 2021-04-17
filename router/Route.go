package router

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"
)

const (
	ROUTE_PREFIX        = "prefix"
	ROUTE_METHOD        = "method"
	ROUTE_RELATIVE_PATH = "relativePath"
	ROUTE_HANDLE        = "handle"
	ROUTE_FLAG          = "flag"
	ROUTE_FRONTPATH     = "frontPath"
	ROUTE_IS_STATIC     = "isStatic"
	ROUTE_DESC          = "desc"
	ROUTE_MIDDLEWARE    = "middleware"
	ROUTE_GROUP_MIDDLE  = "groupMiddle"
	GET                 = "GET"
	POST                = "POST"
	PUT                 = "PUT"
	DELETE              = "DELETE"
	OPTIONS             = "OPTIONS"
	PATCH               = "PATCH"
	HEAD                = "HEAD"
)

var (
	routes             Routes
	unique             *UniqueString
	onlySupportMethods *UniqueString
	once               sync.Once
)

func init() {
	once.Do(func() {
		routes = make(Routes, 0)
		unique = NewUniqueString()
		onlySupportMethods = NewUniqueString()
		onlySupportMethods.Append(GET).
			Append(POST).
			Append(PUT).
			Append(DELETE).
			Append(PATCH).
			Append(OPTIONS).
			Append(HEAD)
	})
}

type UniqueString struct {
	uniSlice []string
	uniMap   map[string]bool
}

func NewUniqueString() *UniqueString {
	return &UniqueString{make([]string, 0), make(map[string]bool)}
}

func (this *UniqueString) Append(key string) *UniqueString {
	this.uniSlice = append(this.uniSlice, key)
	this.uniMap[key] = true
	return this
}

func (this *UniqueString) Exist(key string) bool {
	_, ok := this.uniMap[key]
	return ok
}

func (this *UniqueString) String() string {
	return strings.Join(this.uniSlice, "/")
}

type UniqueStringCallable func(index string, value interface{})

func (this *UniqueString) ForEach(callable UniqueStringCallable) {
	for _, value := range this.uniSlice {
		callable(value, this.uniMap[value])
	}
}

type Route struct {
	groupPrefix  string        // 组前缀
	method       string        // 请求方法 GET/POST/DELETE/PATCH/OPTIONS/HEAD
	relativePath string        // 后端url
	fullPath     string        // 完整url (组前缀 + 后端url)
	handle       interface{}   // 后端控制器函数
	flag         string        // 后端控制器函数标记
	frontPath    string        // 前端 path(前端菜单路由)
	isStatic     bool          // 是否静态文件
	desc         string        // 描述
	middleware   []interface{} // 中间件
	groupMiddle  interface{}   // 组中间件
}

func (this *Route) Prefix() string {
	return this.groupPrefix
}

func (this *Route) Method() string {
	return this.method
}

func (this *Route) RelativePath() string {
	return this.relativePath
}

func (this *Route) Handle() interface{} {
	return this.handle
}

func (this *Route) Flag() string {
	return this.flag
}

func (this *Route) FrontPath() string {
	return this.frontPath
}

func (this *Route) IsStatic() bool {
	return this.isStatic
}

func (this *Route) Desc() string {
	return this.desc
}

func (this *Route) Middleware() interface{} {
	return this.middleware
}

func (this *Route) GroupMiddle() interface{} {
	return this.groupMiddle
}

type Routes []*Route

type RoutesCallable func(index int, route *Route)

func (this *Routes) ForEach(callable RoutesCallable) {
	for key, value := range *this {
		callable(key, value)
	}
}

func AppendRoutes(route *Route) {
	method := strings.ToUpper(route.method)
	if ! onlySupportMethods.Exist(method) {
		panic("route " + route.method + " error, only support:" + onlySupportMethods.String())
	}

	m5 := md5.New()
	m5.Write([]byte(method + ":" + route.fullPath))
	key := hex.EncodeToString(m5.Sum(nil))
	if unique.Exist(key) {
		panic("route " + route.method + ":" + route.fullPath + " already exist")
	}
	unique.Append(key)
	routes = append(routes, route)
}

// 获取路由集
func GetRoutes() *Routes {
	return &routes
}

func Clear() {
	routes = nil
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

func Middleware(value interface{}) *RouteAttribute {
	return NewRouteAttribute(ROUTE_MIDDLEWARE, value)
}

func GroupMiddle(value interface{}) *RouteAttribute {
	return NewRouteAttribute(ROUTE_GROUP_MIDDLE, value)
}
