// 上下文管理

package fgo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// H 简化类型定义
type H map[string]interface{}

type Context struct {
	// 请求-响应对象
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求信息
	Path   string
	Method string
	Params map[string]string
	// 响应信息
	StatusCode int
	// 中间件
	handlers []HandlerFunc
	index    int
	// engine pointer
	engine *Engine
}

// 初始化一个上下文
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	v, _ := c.Params[key]
	return v
}

// PostForm 从表单中获取数据
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 从Query中获取数据
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置Http响应状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 响应类型处理
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		log.Fatal("String write error: ", err)
	}
}

// JSON 类型处理
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 字节流处理
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, err := c.Writer.Write(data)
	if err != nil {
		log.Fatal("Data write error: ", err)
	}
}

// HTML 响应处理
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		log.Println("HTML write error: ", err)
		c.Fail(500, err.Error())
	}
}

// Fail 中间件中断
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}
