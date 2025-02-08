package main

import (
	"fmt"
	"log"
	"net/http"

	"fgo"
)

func main() {
	r := fgo.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			log.Println(err)
		}
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				log.Println(err)
			}
		}
	})

	err := r.Run(":9999")
	if err != nil {
		log.Fatal("服务器运行失败", err)
	}
}
