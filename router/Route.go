package router

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"reflect"
)

type Route struct {
	serve         string        // 服务
	method        string        // 请求方法 GET/POST/DELETE/PATCH/OPTIONS/HEAD
	groupPrefix   string        // 组前缀
	relativePath  string        // 相对路径
	absolutePath  string        // 绝对路径
	handle        interface{}   // 后端控制器函数
	flag          string        // 后端控制器函数标记
	frontPath     string        // 前端 path(前端菜单路由)
	isStatic      bool          // 是否静态文件
	isAuth        bool          // 是否鉴权(默认:false)
	isDataAuth    bool          // 是否数据权限(默认:false)
	isWs          bool          // 是否websocket
	title         string        // 标题
	desc          string        // 描述
	middleware    []interface{} // 中间件
	groupMiddle   []interface{} // 组中间件
	globalMiddle  []interface{} // 全局间件
	handlers      []*Handler
	middlewares   []*Handler
	header        http.Header
	reflectValue  reflect.Value
	funcForPcName string
}

func newRoute() *Route {
	return &Route{handlers: make([]*Handler, 0), middlewares: make([]*Handler, 0)}
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

func (this *Route) AbsolutePath() string {
	return this.absolutePath
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

func (this *Route) IsAuth() bool {
	return this.isAuth
}

func (this *Route) IsDataAuth() bool {
	return this.isDataAuth
}

func (this *Route) IsWs() bool {
	return this.isWs
}

func (this *Route) Title() string {
	return this.title
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

func (this *Route) GlobalMiddle() interface{} {
	return this.globalMiddle
}

func (this *Route) Middlewares() []*Handler {
	return this.middlewares
}

func (this *Route) Handlers() []*Handler {
	return this.handlers
}

func (this *Route) Serve() string {
	return this.serve
}

func (this *Route) Header() http.Header {
	return this.header
}

func (this *Route) SetHeader(header http.Header) *Route {
	this.header = header
	return this
}

func (this *Route) ReflectValue() reflect.Value {
	return this.reflectValue
}

func (this *Route) FuncForPcName() string {
	return this.funcForPcName
}

func UniMd5(method, absolutePath string) string {
	m5 := md5.New()
	m5.Write([]byte(Unique(method, absolutePath)))
	return hex.EncodeToString(m5.Sum(nil))
}

func Unique(method, absolutePath string) string {
	return method + "@" + absolutePath
}
