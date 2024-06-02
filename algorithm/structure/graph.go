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

type GroupInfo struct {
	Parent string
	Times  float64 //倍数
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equMap := make(map[string]GroupInfo)

	//var getParent func(string)string
	getParent := func(a string) string {
		tmp := a
		times := 1.0
		for equMap[a].Parent != a {
			times *= equMap[a].Times
			a = equMap[a].Parent
		}
		equMap[tmp] = GroupInfo{Parent: a, Times: times}
		return a
	}

	for idx, pair := range equations {
		x, y := pair[0], pair[1]
		_, okX := equMap[x]
		_, okY := equMap[y]
		if !okX && !okY { //都不存在

			equMap[x] = GroupInfo{
				Parent: y,
				Times:  values[idx],
			}
			equMap[y] = GroupInfo{
				Parent: y,
				Times:  1.0,
			}
		} else if okX && !okY { //x存在，y不存在,x/y = m,y = x/m
			equMap[y] = GroupInfo{
				Parent: x,
				Times:  1.0 / values[idx],
			}
		} else if okY && !okX { //x不存在，y存在 x/y = m,x = m*y
			equMap[x] = GroupInfo{
				Parent: y,
				Times:  values[idx],
			}
		} else { //都存在,x/y = m, xtimes*ngroupy/(ytimes*groupy) = m,
			xParent := getParent(x) // x = timesx*xparent, y=timesy*yparent,
			yParent := getParent(y)
			if xParent != yParent {
				equMap[xParent] = GroupInfo{Parent: yParent, Times: values[idx] * equMap[y].Times / equMap[x].Times}
			}
		}
	}

	var res []float64
	for _, vals := range queries {
		x, y := vals[0], vals[1]
		_, ok := equMap[x]
		if !ok {
			res = append(res, -1.0)
			continue
		}
		_, ok = equMap[y]
		if !ok {
			res = append(res, -1.0)
			continue
		}
		xparent := getParent(x)
		yparent := getParent(y)
		if xparent != yparent {
			res = append(res, -1.0)
		} else {
			res = append(res, equMap[x].Times/equMap[y].Times)

		}
	}
	return res
}
