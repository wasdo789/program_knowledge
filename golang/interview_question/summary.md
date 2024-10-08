## golang 零碎知识

### 参考资料
[go语言高级编程](https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch1-basic/ch1-03-array-string-and-slice.html)



### 数组
Go语言中数组是值语义。一个数组变量即表示整个数组，它并不是隐式的指向第一个元素的指针（比如C语言的数
组），而是一个完整的值。当一个数组变量被赋值或者被传递的时候，实际上会复制整个数组

var times [5][0]int占用的空间大小是0，可以使用unsafe.Sizeof(times)计算。除此之外无类型的匿名结构体大小也是0，struct{}

### 字符串

底层数据结构是字节数组加字节数组长度（注意不是字符串长度）

```
type StringHeader struct {
    Data uintptr
    Len  int
}
```
因此字符串赋值操作不涉及底层字节数组的改动

- 字符串支持切片 s:="hello world", s1:=s2[:5]
- for range 迭代字符串时，访问的是UTF8解码后的Unicode码点值，比如 for c in range "hello世界"，依次打印：h e l l o 世 界
- len对字符串使用时，计算的事底层字节数组的长度，比如len(hello世界) 计算出的不是7，而是11，（一个utf8中文字符占用3个字节）

### 切片
底层数据结构 reflect.SliceHeader
```
type SliceHeader struct {
    Data uintptr
    Len int
    Cap int
}
```
切片的类型和长度信息无关，只要是相同类型元素构成的切片均对应相同的切片类型

切片可以和nil进行比较，只有当切片底层数据指针为空时切片本身为nil，这时候切片的长度和容量信息将是无效的。如果有切片的底层数据指针为空，但是长度和容量不为0的情况，那么说明切片本身已经被损坏了

- 在容量不足的情况下，append的操作会导致重新分配内存
- 在开头append一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
```
var a = []int{1,2,3}
a = append([]int{0}, a...) //在开头添加1个元素
```
- 切片扩容机制，待补充

### 函数、方法、接口
- 方法是绑定在类型上的，而不是对象
- golang不支持重载
- 方法可以使用
- 使用方法表达式的特性可以将方法还原为普通类型的函数

### 指针
golang里的指针类型
- 普通指针类型
- unsafe.Pointer
- uintptr
区别和联系

### struct 内存对齐

参考 [内存对齐](https://juejin.cn/post/7077833959047954463)
- 重点是unsafe.Alignof(x)的理解。该函数返回值m表示，当变量进行内存对齐时，需要保证x的地址分配到能整除m的地方。对于基础类型，struct都是如何计算的
- struct类型变量大小如何计算的，需要结合各个字段的内存对齐和整体的结构对齐
- 空结构体Struct{} 作为另一个struct的成员变量，在前中后的区别



### make和new的区别
make 只用于map、slice、channel，即引用类型，因为必须初始化，返回对应类型
new 用于任何类型，返回对应指针类型，只分配内存，不初始化，用0填充内存值

new slice时，返回的指针等于nil

### nil的深入理解

### 逃逸分析
#### 闭包捕变量，返回值、参数

### GC 内存回收

### GMP 模型

### sync
#### sync.map
#### sync.Cond
#### Pool,Once

### 性能测试、优化
#### benchmark test 基准测试
#### unit test 单元测试
#### fuzz test

### 零值可用

### 字符串拼接
- +
- fmt.Sprintf()
- strings.Builder
- bytes.Buffer
- []byte



