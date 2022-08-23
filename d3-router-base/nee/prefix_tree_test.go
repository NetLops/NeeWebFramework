package nee

import (
	"fmt"
	"testing"
)

func TestPrefixTreeInsert(t *testing.T) {
	p := prefixTreeNode{}
	p.insert("/p/:lang/name", []string{"p", ":lange", "name"}, 0)
	fmt.Println(p.children[0])
	fmt.Println(p.search([]string{"p", ":lange", "name"}, 0))
	fmt.Println(p.search([]string{"p", ":lange"}, 0))
}
