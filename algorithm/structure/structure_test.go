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
	sts := []*Stu{s1, s2, s3, s4, s5}

	un := NewStuUnion(sts)
	t.Log(un.IsSame(s1, s2))
	t.Log(un.Merge(s1, s2))
	t.Log(un.IsSame(s1, s2))

	t.Log(un.IsSame(s2, s3))
	t.Log(un.Merge(s2, s3))
	t.Log(un.IsSame(s1, s3))
	fmt.Printf("%#v\n", sts)
	fmt.Printf("%#v\n", *un)

}

func TestBFS(t *testing.T) {
	v1 := &Vertex{Num: 1}
	v2 := &Vertex{Num: 2}
	v3 := &Vertex{Num: 3}
	v4 := &Vertex{Num: 4}
	v5 := &Vertex{Num: 5}
	v1.NextNodes = []*Vertex{v3, v2, v5}
	v2.NextNodes = []*Vertex{v3}
	v3.NextNodes = []*Vertex{v4}
	v4.NextNodes = []*Vertex{v2}
	// e1 := &Edge{From: v1, To: v2}
	// e2 := &Edge{From: v1, To: v3}
	// e3 := &Edge{From: v2, To: v3}
	// e4 := &Edge{From: v4, To: v2}
	// e5 := &Edge{From: v3, To: v4}
	//v1.NextEdges = []*Edge{e1, e2}
	// v2.NextEdges = []*Edge{e3}
	// v3.NextEdges = []*Edge{e5}
	// v4.NextEdges = []*Edge{e4}
	graphBFS(v1)
	graphDFS(v1)
}
