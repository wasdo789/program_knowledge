package main

import (
	"fmt"
)

// var wg sync.WaitGroup
func creatEnc() []func(int) int {
	res := []func(int) int{}
	for i := 0; i < 5; i++ {
		f1 := func(int) int {
			return i
		}
		res = append(res, f1)
	}
	return res
}

type Sa struct {
	name string
}

func (s Sa) Name() {
	fmt.Println(s.name)
}

func name(s Sa) {
	fmt.Println(s.name)
}
func A() {
	a, b := 1, 2
	defer func(b int) {
		fmt.Printf("defer a: %d\n", a)
	}(b)
	a = a + b
	b = 10
	fmt.Printf("out a: %d\n", a)
}
func main() {
	var s3 []int
	s1 := new([]int)
	s2 := make([]int, 0)
	fmt.Print(*s1 == nil, s2 == nil)
	fmt.Printf("%+v\n%+v\n%+v\n", s1, s2, s3)

	s4 := []int{1, 2, 3}
	s5 := s4
	fmt.Println(s4)
	fmt.Println(s5)
	s4 = nil
	fmt.Println(s4)
	fmt.Println(s5)
	//A()
	// t1 := reflect.TypeOf(name)
	// t2 := reflect.TypeOf(Sa.Name)
	// fmt.Println(t1, t2, t1 == t2)
	// fs := creatEnc()
	// for _, f := range fs {
	// 	fmt.Println(f())
	// }
	//fmt.Println(runtime.NumCPU())
	// wg.Add(1)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	wg.Wait() // 阻塞在此
	// }()
	// wg.Wait() // 阻塞在此
}
