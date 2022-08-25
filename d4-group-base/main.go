package main

import (
	"fmt"
	"nee"
	"net/http"
)

func main() {
	r := nee.New()
	r.GET("/index", func(c *nee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *nee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Nee</h1>")
		})
		v1.GET("/hello", func(c *nee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you`re at %s\n",
				c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *nee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you`re at %s\n",
				c.Query("name"), c.Path)
		})
		v2.POST("/login", func(c *nee.Context) {
			c.JSON(http.StatusOK,
				nee.H{
					"username": c.PostForm("username"),
					"password": c.PostForm("password"),
				})
		})
	}
	fmt.Println("http://127.0.0.1:9999")
	err := r.Run(":9999")
	if err != nil {
		return
	}
}
