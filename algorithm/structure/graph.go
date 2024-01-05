package structure

import (
	"container/list"
	"fmt"
)

type Vertex struct {
	NextNodes []*Vertex //指向的顶点
	NextEdges []*Edge   //出边
	Num       int       //编号
}

type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
}

type Graph struct {
	Vertexes map[int]*Vertex
}

// BFS,从顶点v开始。
func graphBFS(v *Vertex) {
	if v == nil {
		return
	}
	searched := make(map[*Vertex]bool)
	q := list.New()
	q.PushBack(v)
	for q.Len() > 0 {
		f := q.Remove(q.Front()).(*Vertex)
		fmt.Printf("%v ", f.Num)
		for _, nextVertex := range f.NextNodes {
			_, ok := searched[nextVertex]
			if !ok {
				q.PushBack(nextVertex)
				searched[nextVertex] = true
			}

		}
	}
	fmt.Println()
}

// DFS
func graphDFS(v *Vertex) {
	if v == nil {
		return
	}
	searched := make(map[*Vertex]bool)
	s := list.New()
	s.PushBack(v)
	searched[v] = true
	fmt.Printf("%v ", v.Num)

	for s.Len() > 0 {
		t := s.Back().Value.(*Vertex)
		isDelete := true
		for _, nextNode := range t.NextNodes {
			if _, ok := searched[nextNode]; !ok {
				fmt.Printf("%v ", nextNode.Num)
				s.PushBack(nextNode)
				isDelete = false
				searched[nextNode] = true
				break
			}
		}
		if isDelete {
			s.Remove(s.Back())
		}
	}
	fmt.Println()
}

// 拓扑排序
//func
