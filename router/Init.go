package router

import "sync"

const (
	DefaultServe        = "http"
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
	ROUTE_SERVE         = "serve"
	GET                 = "GET"
	POST                = "POST"
	PUT                 = "PUT"
	DELETE              = "DELETE"
	OPTIONS             = "OPTIONS"
	PATCH               = "PATCH"
	HEAD                = "HEAD"
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
