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




