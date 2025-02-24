package fgo

import (
	"fmt"
	"log"
	"net/http"
)

// HandleFunc 用于定义一个路由的处理函数
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine 用于处理所有的http请求
type Engine struct {
	router map[string]HandleFunc // 存储路由和处理函数的映射关系
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		if err != nil {
			log.Println("ERROR:", err)
			return
		}
	}
}

// New 初始化处理实例
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// 添加路由和处理函数，为了应对同名但不同方法的接口，key用两者拼接
func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// GET 添加GET类型请求
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST 添加POST类型请求
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run 启动服务
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
