package structure

// 单调栈应用
// 计算该右侧首个大于的位置
func rightBigger(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	res := make([]int, len(a))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	ms := make([]int, 0) //slice 模拟栈
	ms = append(ms, 0)
	for i := 1; i < len(a); i++ {
		if len(ms) == 0 || a[i] < a[ms[len(ms)-1]] {
			ms = append(ms, i)
		} else {
			for len(ms) > 0 && a[i] > a[ms[len(ms)-1]] {
				idx := ms[len(ms)-1]
				res[idx] = i
				ms = ms[:len(ms)-1]
			}
			if len(ms) == 0 {
				ms = append(ms, i)
			}
		}
	}
	return res

}
