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
