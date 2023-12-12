package structure

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	r := NewTrie()
	r.Insert("ab")
	r.Insert("abc")
	r.Insert("bc")
	r.Insert("你好")
	fmt.Printf("%#v\n", r)
	fmt.Printf("ab search %d\n", r.Search("ab"))
	fmt.Printf("ab include %d\n", r.PrefixSearch("ab"))
	fmt.Printf("你好 include %d\n", r.PrefixSearch("你好"))
	r.Delete("ab")
	fmt.Printf("ab search %d\n", r.Search("ab"))
	fmt.Printf("ab include %d\n", r.PrefixSearch("ab"))
}
