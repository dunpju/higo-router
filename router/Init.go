package router

import "sync"

const (
	DefaultServe      = "http"
	RoutePrefix       = "prefix"
	RouteMethod       = "method"
	RouteRelativePath = "relativePath"
	RouteHandle       = "handle"
	RouteFlag         = "flag"
	RouteFrontpath    = "frontPath"
	RouteIsStatic     = "isStatic"
	RouteTitle        = "title"
	RouteDesc         = "desc"
	RouteIsAuth       = "isAuth"
	RouteIsDataAuth   = "isDataAuth"
	RouteIsWs         = "isWs"
	RouteMiddleware   = "middleware"
	RouteGroupMiddle  = "groupMiddle"
	RouteServe        = "serve"
	RouteHeader       = "header"
	GET               = "GET"
	POST              = "POST"
	PUT               = "PUT"
	DELETE            = "DELETE"
	OPTIONS           = "OPTIONS"
	PATCH             = "PATCH"
	HEAD              = "HEAD"
	WEBSOCKET         = GET
)

var (
	serve              *Serve
	onlySupportMethods *UniqueString
	once               sync.Once
)

func init() {
	once.Do(func() {
		serve = NewServe(DefaultServe)
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
