// 路由管理

package fgo

import (
	"log"
	"net/http"
)

// 路由管理类
type router struct {
	handlers map[string]HandlerFunc
}

// 初始化路由管理类
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 选择路由处理函数
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
