package nee

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	c := Context{}
	c.Writer = &httptest.ResponseRecorder{}
	header := c.Writer.Header()
	c.Writer.Header().Set("test", "test")
	fmt.Println(c.Writer.Header())
	header.Set("niu", "pi")
	fmt.Println(c.Writer.Header())

	fmt.Printf("%p, %p\n", c.Writer.Header(), header)
	fmt.Println(&header)
}
