package nee

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/asserts/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})

	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")
	if n == nil {
		t.Fatal("nil should`t be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Println(r.getRoutes("GET"))
	fmt.Printf("match path: %s,params['name']: %s\n", n.pattern, ps["name"])
}

func TestUrl(t *testing.T) {
	prefix := http.StripPrefix("/url/go/test", http.FileServer(http.Dir("/Users/netlops")))
	fmt.Println(prefix)
}
