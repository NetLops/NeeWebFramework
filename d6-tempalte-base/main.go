package main

import (
	"fmt"
	"html/template"
	"nee"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormateAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := nee.New()
	r.Use(nee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormateAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *nee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *nee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", nee.H{
			"title":  "nee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *nee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", nee.H{
			"title": "age",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run(":9999")
}
