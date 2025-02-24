package middleware

import (
	"fgo"
	"log"
	"time"
)

func OnlyForV2() fgo.HandlerFunc {
	return func(c *fgo.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
