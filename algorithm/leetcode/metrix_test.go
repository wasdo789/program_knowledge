package leetcode

import "testing"

func TestRotateMetrix(t *testing.T) {
	d := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 16},
		{7, 8, 9, 19},
		{10, 11, 12, 22},
	}
	PrintMetrix(d)
	RotateMetrix(d)
	PrintMetrix(d)

}
