package leetcode

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	total := m * n
	id := 0
	jd := 1
	curCount := 0
	i := 0
	j := -1
	res := []int{}
	for cnt := 0; cnt < total; cnt++ {
		if jd == 1 && curCount == n {
			id = 1
			jd = 0
			m--
			curCount = 0
		} else if id == 1 && curCount == m {
			id = 0
			jd = -1
			n--
			curCount = 0
		} else if jd == -1 && curCount == n {
			id = -1
			jd = 0
			m--
			curCount = 0
		} else if id == -1 && curCount == m {
			id = 0
			jd = 1
			n--
			curCount = 0
		}
		i += id
		j += jd
		res = append(res, matrix[i][j])
		curCount++

	}
	return res
}
