package router

import "fmt"

var (
	currentGroupPrefix string
)

func AddRoute(httpMethod string, route string, handler interface{}) {
	fmt.Println("AddRoute:"+currentGroupPrefix)
}

func AddGroup(prefix string, callable interface{}) {
	previousGroupPrefix := currentGroupPrefix
	currentGroupPrefix += previousGroupPrefix + "/" + prefix
	if _, ok := callable.(func()); ok {
		fmt.Println(currentGroupPrefix)
		fmt.Println(11)
		callable.(func())() // 执行回调
	}
	currentGroupPrefix = previousGroupPrefix
}
