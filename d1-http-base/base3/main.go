package main

import (
	"fmt"
	"nee"
	"net/http"
)

func main() {
	r := nee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL Path = %q\n", req.URL.Path)
		if err != nil {
			fmt.Println("An error has occurred:", err)
			return
		}
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				fmt.Println("An error has occurred:", err)
				return
			}
		}

	})
	if err := r.Run(":9999"); err != nil {
		fmt.Println("Web Server has a error:", err)
		return
	}
}
