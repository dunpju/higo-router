package main

import "github.com/dengpju/higo-router/router"

func main()  {
	router.AddGroup("t1", func() {
		router.AddRoute("GET","tt1", "tthand")
		router.AddGroup("t2", func() {
			router.AddRoute("GET","t21", "tthand")
			router.AddGroup("t3", func() {

			})
			router.AddRoute("GET","t22", "tthand")
		})
		router.AddRoute("GET","tt2", "tthand")
	})

	//router.AddGroup("y1", func() {
	//	router.AddGroup("y2", func() {
	//		router.AddGroup("y3", func() {
	//
	//		})
	//	})
	//})
}
