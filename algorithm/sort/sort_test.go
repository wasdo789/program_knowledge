package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  []int
	result []int
}

func TestSort(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			input:  []int{},
			result: []int{},
		},
		TestCase{
			input:  []int{1},
			result: []int{1},
		},
		TestCase{
			input:  []int{2, 1},
			result: []int{1, 2},
		},
		TestCase{
			input:  []int{2, 1, 3},
			result: []int{1, 2, 3},
		},
		TestCase{
			input:  []int{2, 1, 3, 4},
			result: []int{1, 2, 3, 4},
		},
		TestCase{
			input:  []int{2, 1, 3, 4, 4, 4, 4, 4, 7, 5, 6},
			result: []int{1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7},
		},
	}
	for _, c := range testCases {
		quick_sort(c.input)
		assert.EqualValues(t, c.input, c.result)

	}

}
