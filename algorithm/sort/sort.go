package sort

// 选择排序

// 插入排序

// 快速排序
func partition(d []int, low int, high int) (int, int) {

	if low > high {
		return -1, -1
	} else if low == high {
		return low, low
	}
	less := low - 1 //小于区右边界,闭区间
	more := high    //大于区左边界，闭区间
	idx := low      //扫描位置
	val := d[high]
	for idx < more {
		//遇到小于val的，idx和less+1位置互换，less++，idx++
		if d[idx] < val {
			d[less+1], d[idx] = d[idx], d[less+1]
			less++
			idx++
		} else if d[idx] == val {
			//遇到等于val的，idx++
			idx++
		} else {
			//遇到大于val的，和more-1位置互换，more--
			d[more-1], d[idx] = d[idx], d[more-1]
			more--
		}
	}
	//high和more-1位置互换,返回less+1，more-1
	d[high], d[more] = d[more], d[high]
	return less + 1, more
}

func sort_range(d []int, low int, high int) {
	if low >= high {
		return
	}

	// idx := rand.Intn(high-low) + low
	// d[idx],d[hihigh] = d[high],d[idx]
	begin, end := partition(d, low, high)
	sort_range(d, low, begin-1)
	sort_range(d, end+1, high)
}

func quick_sort(datas []int) {
	sort_range(datas, 0, len(datas)-1)
}

// 堆排序
// 逐步建立堆，自底向上
func push_heap(datas []int, heapSize int) {
	//一直和父节点对比
	if heapSize <= 0 {
		return
	}
	idx := heapSize
	for idx > 0 && datas[idx] > datas[(idx-1)/2] {
		datas[idx], datas[(idx-1)/2] = datas[(idx-1)/2], datas[idx]
		idx = (idx - 1) / 2
	}
}

// 把数组初始化为堆，自底向上

// 调整以下表j为堆顶的堆
func heapify(datas []int, j int, heapSize int) {
	for 2*j+1 < heapSize {
		larggerIdx := 2*j + 1
		if 2*j+2 < heapSize && datas[2*j+2] > datas[2*j+1] {
			larggerIdx = 2*j + 2
		}
		if datas[larggerIdx] <= datas[j] {
			break
		}
		datas[j], datas[larggerIdx] = datas[larggerIdx], datas[j]
		j = larggerIdx
	}
}
func init_heap(datas []int) {
	if len(datas) <= 1 {
		return
	}
	//只调整非叶子节点
	for i := len(datas)/2 - 1; i >= 0; i-- {
		heapify(datas, i, len(datas))
	}
}

// 删除堆顶数据后，重新调整，自顶向下
func pop_heap(datas []int, heapSize int) {
	datas[0], datas[heapSize-1] = datas[heapSize-1], datas[0]
	heapSize--
	idx := 0
	heapify(datas, idx, heapSize)
}
func heap_sort(datas []int) {
	//if len(datas)
	// //push heap
	// for idx, _ := range datas {
	// 	push_heap(datas, idx)
	// }
	init_heap(datas)
	//pop
	for idx, _ := range datas {
		pop_heap(datas, len(datas)-idx)
	}
}
