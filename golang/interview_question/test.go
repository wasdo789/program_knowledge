package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func TestArray(a [3]int) {
	a[0] = 1
	a[1] = 1
	a[2] = 1
}

func TestArrayPointer(a *([3]int)) {
	a[0] = 1
	a[1] = 1
	a[2] = 1
}

func main() {
	fmt.Println("//测试数组是值语义")
	b := [...]int{1, 2, 3}
	fmt.Println(b)
	TestArray(b)
	fmt.Println(b)

	fmt.Println("//测试数组指针")
	c := &b
	fmt.Println(b)
	TestArrayPointer(c)
	fmt.Println(b)

	fmt.Println("//空数组大小为0")
	var times [5][0]int
	fmt.Printf("size of times %d", unsafe.Sizeof(times))

	fmt.Println("字符串测试")
	s := "hello世界"
	fmt.Printf("%#v\n", []byte(s))
	fmt.Printf("len(\"hello世界\") is %d\n", len(s))
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c)
	}
	fmt.Printf("[]rune(\"hello世界\")=%#v", []rune("hello世界"))

	fmt.Println("使用指针强制更改字符串底层数据")
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytes := []byte("abc")
	hdr.Data = (uintptr)(unsafe.Pointer(&bytes[0]))
	hdr.Len = len(bytes)
	fmt.Println(s)

}
