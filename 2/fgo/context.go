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
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

// 初始化一个上下文
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
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

// String响应类型处理
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
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		log.Fatal("HTML write error: ", err)
	}
}
