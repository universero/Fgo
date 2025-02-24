package main

import (
	"log"
	"net/http"

	"fgo"
)

func main() {
	r := fgo.New()
	r.GET("/", func(c *fgo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *fgo.Context) {
		// expect /hello?name=fgo
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *fgo.Context) {
		// expect /hello/fgo
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.POST("/login", func(c *fgo.Context) {
		c.JSON(http.StatusOK, fgo.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.GET("/assets/*filepath", func(c *fgo.Context) {
		c.JSON(http.StatusOK, fgo.H{"filepath": c.Param("filepath")})
	})

	err := r.Run(":9999")
	if err != nil {
		log.Fatal(err)
	}
}
