package router

import (
	"reflect"
	"runtime"
)

type Handler struct {
	handlerFunc   interface{}
	reflectValue  reflect.Value
	funcForPcName string
}

func newHandler(handlerFunc interface{}) *Handler {
	handler := &Handler{handlerFunc: handlerFunc}
	switch handlerFunc.(type) {
	case string, int, int64, int32, int8, int16:
	default:
		handler.reflectValue = reflect.ValueOf(handlerFunc)
		handler.funcForPcName = runtime.FuncForPC(handler.reflectValue.Pointer()).Name()
	}
	return handler
}

func (m *Handler) HandlerFunc() interface{} {
	return m.handlerFunc
}

func (m *Handler) ReflectValue() reflect.Value {
	return m.reflectValue
}

func (m *Handler) FuncForPcName() string {
	return m.funcForPcName
}
