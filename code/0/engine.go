// 通过自定义实例处理请求

package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 用于处理所有的http请求
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		log.Println("fgo engine is dealing ", req.URL.Path)
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			log.Println(err)
		}
	case "/hello":
		log.Println("fgo engine is dealing ", req.URL.Path)
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				log.Println(err)
			}
		}
	default:
		log.Println("fgo engine is dealing ", req.URL.Path)
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	engine := new(Engine)
	log.Println(http.ListenAndServe(":9999", engine))
}
