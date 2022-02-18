package router

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

type Route struct {
	serve        string        // 服务
	method       string        // 请求方法 GET/POST/DELETE/PATCH/OPTIONS/HEAD
	groupPrefix  string        // 组前缀
	relativePath string        // 相对路径
	absolutePath string        // 绝对路径
	handle       interface{}   // 后端控制器函数
	flag         string        // 后端控制器函数标记
	frontPath    string        // 前端 path(前端菜单路由)
	isStatic     bool          // 是否静态文件
	isAuth       bool          // 是否鉴权(默认:true)
	isWs         bool          // 是否websocket
	desc         string        // 描述
	middleware   []interface{} // 中间件
	groupMiddle  []interface{} // 组中间件
	unimd5       string        // 唯一标识 md5(method + "@" + absolutePath)
	unique       string        // 唯一标识 method + "@" + absolutePath
	header       http.Header
}

func newRoute() *Route {
	return &Route{}
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

func (this *Route) IsWs() bool {
	return this.isWs
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

func (this *Route) Serve() string {
	return this.serve
}

func (this *Route) GenUniMd5() *Route {
	this.unimd5 = UniMd5(this.method, this.absolutePath)
	this.unique = Unique(this.method, this.absolutePath)
	return this
}

func (this *Route) UniMd5() string {
	return this.unimd5
}

func (this *Route) Unique() string {
	return this.unique
}

func (this *Route) Header() http.Header {
	return this.header
}

func (this *Route) SetHeader(header http.Header) *Route {
	this.header = header
	return this
}

func UniMd5(method, absolutePath string) string {
	m5 := md5.New()
	m5.Write([]byte(Unique(method, absolutePath)))
	return hex.EncodeToString(m5.Sum(nil))
}

func Unique(method, absolutePath string) string {
	return method + "@" + absolutePath
}
