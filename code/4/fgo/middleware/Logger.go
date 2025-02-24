package middleware

import (
	"fgo"
	"log"
	"time"
)

func Logger() fgo.HandlerFunc {
	return func(c *fgo.Context) {
		// 开始时间
		start := time.Now()
		// 执行请求
		c.Next()
		// 结束时间
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(start))
	}
}
