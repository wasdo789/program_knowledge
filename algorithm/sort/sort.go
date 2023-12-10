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
func push_heap(datas []int, val int, heapSize int) {
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
func pop_heap(datas []int, val int, heapSize int) {
	datas[0], datas[heapSize-1] = datas[heapSize-1], datas[0]
	heapSize--
	idx := 0
	for heapSize > 1 && idx <= (heapSize-2)/2 {
		larggerIdx := idx*2 + 1
		if idx*2+2 < heapSize && datas[idx*2+2] > datas[idx*2+1] {
			larggerIdx = idx*2 + 2
		}
		if datas[larggerIdx] <= datas[idx] {
			break
		}
		datas[larggerIdx], datas[idx] = datas[idx], datas[larggerIdx]
		idx = larggerIdx
	}
}
func heap_sort(datas []int) {
	//if len(datas)
	//push heap
	for idx, v := range datas {
		push_heap(datas, v, idx)
	}
	//pop
	for idx, v := range datas {
		pop_heap(datas, v, len(datas)-idx)
	}
}
