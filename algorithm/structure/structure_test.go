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

func TestStuUnion(t *testing.T) {
	s1 := &Stu{Name: "aa"}
	s2 := &Stu{Name: "bb"}
	s3 := &Stu{Name: "c"}
	s4 := &Stu{Name: "d"}
	s5 := &Stu{Name: "e"}
	un := NewStuUnion([]*Stu{s1, s2, s3, s4, s5})
	t.Log(un.IsSame(s1, s2))
	t.Log(un.Merge(s1, s2))
	t.Log(un.IsSame(s1, s2))

	t.Log(un.IsSame(s1, s3))
	t.Log(un.Merge(s3, s2))
	t.Log(un.IsSame(s1, s3))

}
