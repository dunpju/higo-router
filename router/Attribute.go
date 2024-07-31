package router

import (
	"net/http"
)

type RouteAttributes []*RouteAttribute

func (this RouteAttributes) Find(name string) interface{} {
	for _, p := range this {
		if p.Name == name {
			if p.Name != RouteMiddleware && p.Name != RouteGroupMiddle && p.Name != RouteGlobalMiddle {
				return p.Value[0]
			} else {
				return p.Value
			}
		}
	}
	return nil
}

func (this RouteAttributes) Append(attribute *RouteAttribute) RouteAttributes {
	this = append(this, attribute)
	return this
}

type RouteAttribute struct {
	Name  string
	Value []interface{}
}

func NewRouteAttribute(name string, values ...interface{}) *RouteAttribute {
	if len(values) == 0 {
		panic("route attribute " + name + " can't be empty")
	}
	return &RouteAttribute{Name: name, Value: values}
}

func Flag(value string) *RouteAttribute {
	return NewRouteAttribute(RouteFlag, value)
}

func FrontPath(value string) *RouteAttribute {
	return NewRouteAttribute(RouteFrontpath, value)
}

func IsStatic(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteIsStatic, value)
}

func Title(value string) *RouteAttribute {
	return NewRouteAttribute(RouteTitle, value)
}

func Desc(value string) *RouteAttribute {
	return NewRouteAttribute(RouteDesc, value)
}

func IsAuth(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteIsAuth, value)
}

func IsDataAuth(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteIsDataAuth, value)
}

func CancelGlobalGroupPrefix(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteCancelGlobalGroupPrefix, value)
}

func CancelGlobalApiGroupPrefix(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteCancelGlobalApiGroupPrefix, value)
}

func IsWs(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteIsWs, value)
}

func Middleware(value ...interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteMiddleware, value...)
}

func GroupMiddle(value ...interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteGroupMiddle, value...)
}

func GlobalMiddle(value ...interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteGlobalMiddle, value...)
}

func SetServe(value interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteServe, value)
}

func SetHeader(value http.Header) *RouteAttribute {
	return NewRouteAttribute(RouteHeader, value)
}
