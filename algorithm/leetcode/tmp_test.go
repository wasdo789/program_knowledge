package leetcode

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(n, 3)
	t.Log(n)
}

func TestCanCompleteCircuit(t *testing.T) {
	t.Log(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}

func TestIntToRoman(t *testing.T) {
	t.Log(intToRoman(1994))
}

func TestReverseWords(t *testing.T) {
	t.Log(reverseWords("  hello  world  "))

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[:3]
	fmt.Printf("s1 len:%d, cap:%d\n", len(s1), cap(s1))
	s3 := arr[:3:3]
	fmt.Printf("s3 len:%d, cap:%d\n", len(s3), cap(s3))
	s2 := arr[3:5]
	fmt.Printf("s2 len:%d, cap:%d\n", len(s2), cap(s2))
	s1 = append(s1, 33)
	fmt.Println(s1)
	fmt.Println(arr)
	a := make([]int, 0, 0)
	fmt.Printf("len:%d, cap:%d\n", len(a), cap(a))
	for i := 1; i <= 1000; i++ {
		a = append(a, i)
		fmt.Printf("len:%d, cap:%d\n", len(a), cap(a))
	}

}

func TestMaxArea(t *testing.T) {
	//
	//t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	//t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//t.Log(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	//t.Log(findSubstring("aaaaaaaaaaaaaa", []string{"aa", "aa"}))
	//t.Log(minWindow("ADOBECODEBANC", "ABC"))
	// t.Log(minWindow("babb", "baba"))
	// t.Log(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	// t.Log(maxSubArray([]int{8, -19, 5, -4, 20}))
	// t.Log(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// t.Log(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// s := []int{1}
	// t.Log(s[2:])
	r := buildTree([]int{3, 1, 2, 4}, []int{1, 2, 3, 4})
	t.Log(r)
}
