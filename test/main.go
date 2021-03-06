package main

import (
	"github.com/dengpju/higo-router/router"
)

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

	router.AddGroup("/y1", func() {
		router.AddRoute("GET","/y1-r1", "y1-r1-hand")
		router.AddGroup("/y2", func() {
			router.AddRoute("GET","/y2-r1", "y2-r1-hand")
			router.AddGroup("/y3", func() {
				router.AddRoute("GET","/y3-r1", "y3-r1-hand")
				router.AddRoute("GET","/y3-r2", "y3-r2-hand")
				router.Get("/get_test", "get_test")
				router.Post("/post_test", "post_test")
				router.Put("/put_test", "put_test")
				router.Delete("/delete_test", "delete_test")
				router.Patch("/patch_test", "patch_test")
				router.Head("/head_test", "head_test")
			})
			router.AddRoute("GET","/y2-r2", "y2-r2-hand")
		})
		router.AddRoute("GET","/y1-r2", "y1-r2-hand")
	})

	router.AddRoute("GET","/x1-r1", "x1-r1-hand")

}
