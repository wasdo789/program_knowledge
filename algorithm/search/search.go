package search

func combine(n int, k int) [][]int {
	path := make([]int, 0, k)
	var res [][]int
	combineRecur(1, n, k, &path, &res)
	return res
}

func combineRecur(beginIdx int, n int, k int, path *([]int), res *([][]int)) {
	if len(*path) == k {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
	} else {
		for i := beginIdx; i <= n; i++ {
			*path = append(*path, i)
			combineRecur(i+1, n, k, path, res)
			*path = (*path)[:len(*path)-1]
		}
	}
}
