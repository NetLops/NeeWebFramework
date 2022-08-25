package main

import (
	"fmt"
	"log"
	"nee"
	"net/http"
	"time"
)

func onlyForV2() nee.HandlerFunc {
	return func(c *nee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		fmt.Println(1)
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode,
			c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := nee.New()
	r.Use(nee.Logger()) // global middleware
	r.GET("/", func(c *nee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Nee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *nee.Context) {
			// expect /hello/geektutu
			fmt.Println(2)
			c.String(http.StatusOK,
				"hello %s, you`re at %s\n",
				c.Param("name"), c.Path)

		})
	}

	r.Run(":9999")
}
