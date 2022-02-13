package router

import "net/http"

type RouteAttributes []*RouteAttribute

func (this RouteAttributes) Find(name string) interface{} {
	for _, p := range this {
		if p.Name == name {
			return p.Value
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
	Value interface{}
}

func NewRouteAttribute(name string, value interface{}) *RouteAttribute {
	return &RouteAttribute{Name: name, Value: value}
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

func Desc(value string) *RouteAttribute {
	return NewRouteAttribute(RouteDesc, value)
}

func IsAuth(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteIsAuth, value)
}

func IsWs(value bool) *RouteAttribute {
	return NewRouteAttribute(RouteIsWs, value)
}

func Middleware(value interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteMiddleware, value)
}

func GroupMiddle(value interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteGroupMiddle, value)
}

func SetServe(value interface{}) *RouteAttribute {
	return NewRouteAttribute(RouteServe, value)
}

func SetHeader(value http.Header) *RouteAttribute {
	return NewRouteAttribute(RouteHeader, value)
}
