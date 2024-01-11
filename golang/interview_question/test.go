package main

import (
	"fmt"
	"math"
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

	//指针的示例
	fmt.Println("----------pointer test -----------------")
	un1 := uint(math.MaxUint64)
	n2 := int(3)
	//普通指针，可以用于取值操作
	fmt.Println(reflect.TypeOf(un1))
	fmt.Println(reflect.TypeOf(n2))
	fmt.Println(*(&un1))

	fmt.Println(reflect.TypeOf(&un1))
	fmt.Println(reflect.TypeOf(&n2))

	//pointer用于指针强制转换，不可用*取值操作
	//p1 := (*int)(&un1) //编译错误：cannot convert &un1 (value of type *uint) to *int
	p1 := (*int)(unsafe.Pointer(&un1))
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(*p1)
	pt := unsafe.Pointer(&un1)
	fmt.Println(reflect.TypeOf(pt))
	//fmt.Println(*pt) //出错，invalid operation: cannot indirect pt (variable of type unsafe.Pointer)

	//unptr 1.用于和Pointer互相转换。 2、用于指针运算，比如偏移
	arr := new([3]int)
	p1 = (*int)(unsafe.Pointer(arr))
	//up1 := uintptr(arr) //不能直接把普通指针转为uintptr， cannot convert arr (variable of type *[3]int) to uintptr
	*p1 = 2 //第一个元素改为2
	p2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + uintptr(unsafe.Sizeof(int(0)))))
	*p2 = 3 //第二个元素改为3
	fmt.Println(*arr)

	fmt.Println("----------struct memory align test -----------------")
	//测试内置类型对齐大小，返回值是min（机器位数/8，类型字节大小)
	fmt.Println(unsafe.Alignof(byte(1)))       // 1 -- min(8,1)
	fmt.Println(unsafe.Alignof(int16(1)))      // 1 -- min(8,1)
	fmt.Println(unsafe.Alignof(int32(1)))      // 4 -- min (8,4)
	fmt.Println(unsafe.Alignof(int64(1)))      // 8 -- min (8,8)
	fmt.Println(unsafe.Alignof(complex128(1))) // 8 -- min(8,16)
	//数组、slice、map内存对齐大小和元素数量无关， 和底层结构有关,实际上和struct规则一致
	fmt.Println("unsafe.Alignof([]byte{1, 2}) ", unsafe.Alignof([]byte{1, 2}))

	type Sw struct {
		b byte
		i int32
		j int64
	}
	w := new(Sw)
	//结构体的内存对齐大小为其各个字段内存对齐大小的最大值
	fmt.Printf("sizeof(w) %v, alignof(Sw) %v\n", unsafe.Sizeof(*w), unsafe.Alignof(w))
	//其实alignof只和类型有关，和在不在struct无关
	fmt.Printf("alignof(Sw.b) %v,alignof(Sw.i) %v,alignof(Sw.j) %v\n", unsafe.Alignof(w.b), unsafe.Alignof(w.i), unsafe.Alignof(w.j))

	//改变成员变量位置，结构体大小发生变化
	type Sw2 struct {
		b byte
		j int64
		i int32
	}
	w2 := new(Sw2)
	//结构体
	fmt.Printf("sizeof(w2) %v, alignof(Sw2) %v\n", unsafe.Sizeof(*w2), unsafe.Alignof(w2))
	//空结构体作为成员变量
	type SC struct {
		a struct{}
		b int64
		c int64
	}

	type SD struct {
		a int64
		b struct{}
		c int64
	}

	type SE struct {
		a int64
		b int64
		c struct{}
	}
	sc := SC{}
	sd := SD{}
	se := SE{}

	fmt.Println(unsafe.Sizeof(sc)) // 16
	fmt.Println(unsafe.Sizeof(sd)) // 16
	fmt.Println(unsafe.Sizeof(se)) // 24
	fmt.Printf("%p, %p\n", &sc.a, &sc.b)

	// //make slice容量0和new的区别
	// sl1 := make([]int, 0)
	// fmt.Printf("sl1 %#v\n", sl1)
	// slh1 := (*reflect.SliceHeader)(unsafe.Pointer(&sl1))
	// fmt.Printf("slice header %#v, *s1h1.Data: %v\n", slh1, *(&slh1.Data))

	// sl2 := new([]int)
	// fmt.Printf("sl2 %#v\n", *sl2)
	// slh2 := (*reflect.SliceHeader)(unsafe.Pointer(sl2))
	// fmt.Printf("slice header %#v\n", slh2)
	// *sl2 = append(*sl2, 10)
	// fmt.Printf("sl2 %#v\n", *sl2)
	// fmt.Printf("slice header %#v\n", slh2)
	m1 := make(map[int]int)
	fmt.Println(m1)

}
