package leetcode

import "testing"

func TestRotate(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(n, 3)
	t.Log(n)
}

func TestCanCompleteCircuit(t *testing.T) {
	t.Log(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
