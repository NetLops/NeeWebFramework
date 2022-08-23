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
		// expect /hello?name=geetutu
		c.String(http.StatusOK, "hello %s, you`re at %s\n",
			c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *nee.Context) {
		// expect /hello/geetutu
		c.String(http.StatusOK, "hello %s, you`re at %s\n",
			c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *nee.Context) {
		c.JSON(http.StatusOK, nee.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":9999")
}
