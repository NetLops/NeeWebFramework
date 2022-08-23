package main

import (
	"nee"
	"net/http"
)

func main() {
	r := nee.New()
	r.GET("/", func(c *nee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Nee</h1>")
	})
	r.GET("/hello", func(c *nee.Context) {
		c.String(http.StatusOK, "hello %s, you`re at 5s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *nee.Context) {
		c.JSON(http.StatusOK, nee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
