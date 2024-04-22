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

func rotate2(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
}

func setZeroes(matrix [][]int) {
	col0Flag := false
	row0Flag := false
	c := len(matrix[0])
	r := len(matrix)
	for j := 0; j < c; j++ {
		if matrix[0][j] == 0 {
			row0Flag = true
			break
		}
	}
	for j := 0; j < r; j++ {
		if matrix[j][0] == 0 {
			col0Flag = true
			break
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < c; j++ {
		if row0Flag {
			matrix[0][j] = 0
		}
	}
	for j := 0; j < r; j++ {
		if col0Flag {
			matrix[j][0] = 0
		}
	}
}
