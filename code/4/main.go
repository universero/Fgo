package main

import (
	"fgo/middleware"
	"log"
	"net/http"

	"fgo"
)

func main() {
	r := fgo.New()
	r.Use(middleware.Logger()) // global middleware
	r.GET("/", func(c *fgo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(middleware.OnlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *fgo.Context) {
			// expect /hello/fgo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	err := r.Run(":9999")
	if err != nil {
		log.Fatal(err)
	}
}
