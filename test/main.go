package main

import (
	"fmt"
	"github.com/dengpju/higo-router/router"
	"sync"
)

func Test() {

}

func main() {

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()

	}()

	go func() {
		wg.Add(1)
		defer wg.Done()

	}()
	wg.Wait()
	router.SetInitGroupIsAuth(true)
	router.AddGroup("/t1", func() {
		router.AddRoute("GET", "/t1-r1", "t1-r1-hand", router.IsAuth(false))
		router.AddGroup("/t2", func() {
			router.AddRoute("GET", "/t2-r1", "tt2-r1-hand")
			router.AddGroup("/t3", func() {
				router.AddRoute("GET", "/t3-r1", "t3-r1-hand", router.Middleware(func() {}))
				router.AddRoute("GET", "/t3-r2", "t3-r2-hand")
			}, router.IsAuth(false), router.GroupMiddle(func() {}))
			router.AddRoute("GET", "/t2-r2", "t2-r2-hand")
		}, router.IsAuth(true), router.GroupMiddle(func() {}))
		router.AddRoute("GET", "/t1-r2", "t1-r2-hand")
	}, router.GroupMiddle(func() {}, func() {}, func() {}))

	router.AddGroup("/y1", func() {
		router.AddRoute("GET", "/y1-r1", "y1-r1-hand")
		router.AddGroup("/y2", func() {
			router.AddRoute("GET", "/y2-r1", "y2-r1-hand")
			router.AddGroup("/y3", func() {
				router.AddRoute("GET", "/y3-r1", "y3-r1-hand")
				router.AddRoute("GET", "/y3-r2", "y3-r2-hand")
				router.Get("/get_test", "get_test")
				//router.Get("/get_test", "get_test") // 测试 panic: route GET:/y1/y2/y3/get_test already exist
				router.Post("/post_test", "post_test")
				router.Put("/put_test", "put_test")
				router.Delete("/delete_test", "delete_test")
				router.Patch("/patch_test", "patch_test")
				router.Head("/head_test", "head_test")
				//router.Head("/head_test/:id", "head_test")
				router.Head("/head_test/:id/:name", "head_test")
				router.Head("/head_test/:aa/:bb", "head_test")
				router.Head("/head_test/:id/:name/tt", "head_test")
			})
			router.AddRoute("GET", "/get_test", "y2-r2-hand")
		})
		router.AddRoute("GET", "/y1-r2", "y1-r2-hand")
	})
	router.AddRoute("GET", "/y1-r3", "y1-r3-hand")
	fmt.Println(router.GetServes())
	router.GetRoutes(router.DefaultServe).ForEach(func(index int, route *router.Route) {
		fmt.Println(route)
	})
	router.GetRoutes(router.DefaultServe).Trie().Each(func(n *router.Node) {
		fmt.Println(*n)
	})
	fmt.Println("================")
	n, e := router.GetRoutes(router.DefaultServe).Search(router.HEAD, "/y1/y2/y3/head_test/1/gg/tt")
	if n != nil {
		fmt.Println(n)
		fmt.Println(n.Route)
	} else {
		fmt.Println(n, e)
	}
	fmt.Println("================")
	router.GetRoutes(router.DefaultServe).Trie().Each(func(n *router.Node) {
		fmt.Println(*n)
	})
	fmt.Println("================2222")
	fmt.Println(router.GetRoutes(router.DefaultServe).Search(router.HEAD, "/y1/y1-r21"))
	fmt.Println("================")

	// 增加 serve
	//router.AddServe("https").
	//	AddRoute("GET", "/x1-r1", "x1-r1-hand1", router.Flag("fff"))
	//router.AddServe("http").
	//	AddRoute("GET", "/x1-r11", "x1-r1-hand2")
	//router.AddRoute("GET", "/gggggggggggggggg", "y2-r2-hand")
	//router.AddServe("https").
	//	AddRoute("GET", "/fffffffffff", "x1-r1-hand1").
	//	Ws("/ffffffffff", "ws")
	//
	//fmt.Println(len(router.GetRoutes(router.DefaultServe).List()))
	//router.GetRoutes(router.DefaultServe).ForEach(func(index int, route *router.Route) {
	//	fmt.Println(route)
	//})
	//fmt.Println(router.GetServes())
	//router.GetRoutes("https").ForEach(func(index int, route *router.Route) {
	//	fmt.Println(route)
	//})
	//fmt.Println(router.GetRoutes("http").Route("GET", "/y1/y2/y3/get_test"))

	/**
	fmt.Println(len(*router.GetRoutes()))
	router.AddRoute("GET","/t1-r1", "t1-r1-hand")
	fmt.Println(len(*router.GetRoutes()))


	router.AddRoute("GET","/t1-r1", "t1-r1-hand")
	router.AddRoute("GET","/test1", Test, router.Flag("test"))
	router.AddRoute("GET","/test", Test, router.Flag("test"))
	fmt.Println(*router.GetRoutes())
	router.GetRoutes().ForEach(func(index int, route *router.Route) {
		fmt.Println(route)
	})

	*/
}
