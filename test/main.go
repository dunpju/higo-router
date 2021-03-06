package main

import "github.com/dengpju/higo-router/router"

func main()  {
	router.AddGroup("/t1", func() {
		router.AddRoute("GET","/t1-r1", "t1-r1-hand")
		router.AddGroup("/t2", func() {
			router.AddRoute("GET","/t2-r1", "tt2-r1-hand")
			router.AddGroup("/t3", func() {
				router.AddRoute("GET","/t3-r1", "t3-r1-hand")
			})
			router.AddRoute("GET","/t2-r2", "t2-r2-hand")
		})
		router.AddRoute("GET","/t1-r2", "t1-r2-hand")
	})

	//router.AddGroup("y1", func() {
	//	router.AddGroup("y2", func() {
	//		router.AddGroup("y3", func() {
	//
	//		})
	//	})
	//})
}
