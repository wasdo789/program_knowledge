package leetcode

import "fmt"

// 正方形矩阵旋转90度
func RatateEdge(data [][]int, r1 int, r2 int) {
	for i := 0; i < r2-r1; i++ {
		tmp := data[r1][r1+i]
		data[r1][r1+i] = data[r2-i][r1]
		data[r2-i][r1] = data[r2][r2-i]
		data[r2][r2-i] = data[r1+i][r2]
		data[r1+i][r2] = tmp
	}

}
func RotateMetrix(data [][]int) {
	edge := len(data)
	r1 := 0
	r2 := edge - 1
	for r1 < r2 {
		RatateEdge(data, r1, r2)
		r1++
		r2--
	}
}

func PrintMetrix(data [][]int) {
	edge := len(data)
	for i := 0; i < edge; i++ {
		for j := 0; j < edge; j++ {
			fmt.Printf("%d ", data[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
