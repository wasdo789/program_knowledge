package structure

/*
利用go标准库实现堆和优先队列
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

type TreeNodeHeap []*TreeNode

func (h TreeNodeHeap) Len() int           { return len(h) }
func (h TreeNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h TreeNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TreeNodeHeap) Push(x any) {
	*h = append(*h, x.(*TreeNode))
}
func (h *TreeNodeHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}
