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

	fmt.Println("测试切片")
	s1 := make([]int, 8, 16)
	fmt.Printf("s1 len %d, cap %d\n", len(s), cap(s1))
	//fmt.Printf("access out len, s1[4] %d\n", s1[4]) //访问slice超过len会报数组越界
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("slice header %#v\n", sh)

	fmt.Println("切片操作")
	fmt.Println(s1)
	s1 = append(s1, 1) //尾部插入
	fmt.Println(s1)
	s1 = append(s1[:2], append([]int{8}, s1[2:]...)...) //中间位置2插入8
	fmt.Println(s1)
	//使用copy避免中间slice的方法
	s1 = append(s1, 0)
	copy(s1[3:], s1[2:])
	s1[2] = 88
	fmt.Println(s1)

	s1 = s1[:len(s1)-2] //删除尾部2个元素，没有移动数据指针
	fmt.Println(s1)
	sh = (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("slice header %#v\n", sh)
	s1 = s1[1:] //删除开头1个元素，底层数据指针发生了变化
	sh = (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("slice header %#v\n", sh)
	fmt.Println(s1)
	s1 = append(s1[:0], s1[1:]...) //原地删除开头一个元素，底层数据指针没变化
	fmt.Printf("slice header %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Println(s1)
	//删除slice元素时，cap没发生变化
	pos := 1
	s1 = append(s1[:pos], s1[pos+1:]...) //删除中间位置 pos 元素
	fmt.Println(s1)
	s1 = s1[:pos+copy(s1[pos:], s1[pos+2:])] //使用copy删除中间元素
	fmt.Println(s1)

	s2 := s1[:0]
	fmt.Printf("s2: %#v, s2 == nil, %v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s2)), s2 == nil)
	var s3 []int
	fmt.Printf("s3: %#v, s3 == nil, %v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s3)), s3 == nil)

}
