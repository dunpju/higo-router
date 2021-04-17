package router

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

func SetServe(value interface{}) *RouteAttribute {
	return NewRouteAttribute(ROUTE_SERVE, value)
}
