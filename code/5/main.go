package main

import (
	"fgo"
	"fmt"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := fgo.Default()
	r.GET("/", func(c *fgo.Context) {
		c.String(http.StatusOK, "Hello fgo\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *fgo.Context) {
		names := []string{"fgo"}
		c.String(http.StatusOK, names[100])
	})
	err := r.Run(":9999")
	if err != nil {
		log.Fatal(err)
	}
}
