package structure

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 7, 3, 6, 1}
	heap.Init(h)
	fmt.Print("#v\n", h)
	for len(*h) > 0 {
		//t.Log(h.Pop())
		t.Log(heap.Pop(h))
	}
}

func TestHuffman(t *testing.T) {
	head := Huffman([]int{2, 7, 3, 6, 1})
	PreTraversal(head)
}
