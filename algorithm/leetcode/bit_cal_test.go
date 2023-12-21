package leetcode

import (
	"fmt"
	"testing"
)

func TestLowBitOne(t *testing.T) {
	a := 0
	t.Log(LowBitOne(a))

	a = 1
	t.Log(LowBitOne(a))
	a = 5
	fmt.Printf("%b\n", a)
	t.Log(LowBitOne(a))
	a = -6
	fmt.Printf("%b\n", a)
	t.Log(LowBitOne(a))
}

func TestOneCount(t *testing.T) {
	a := -5
	fmt.Printf("%b\n", uint(a))
	t.Log(OneCount(a))

	a = 5
	fmt.Printf("%b\n", uint(a))
	t.Log(OneCount(a))
}

func TestFindOddTimesNumber(t *testing.T) {
	t.Log(FindOddTimesNumber([]int{1, 1, 2}))
	t.Log(FindOddTimesNumber([]int{1, 2, 2, 1, 3}))
	t.Log(FindOddTimesNumber([]int{1, 2, 2, 1, -3}))
}

func TestFind2OddTimesNumber(t *testing.T) {
	t.Log(Find2OddTimesNumber([]int{2, 3}))
	t.Log(Find2OddTimesNumber([]int{1, 2, 2, 1, 3, 4}))
}
