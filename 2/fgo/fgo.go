package fgo

import (
	"net/http"
)

// HandlerFunc 定义基于Context的处理函数
type HandlerFunc func(*Context)

// Engine 核心类，处理请求
type Engine struct {
	router *router
}

// New 初始化核心类
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 添加GET类型路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加POST类型路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
