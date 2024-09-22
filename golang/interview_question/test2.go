package main

import "fmt"

type T int
func search(nums []int, target int) int {
	len := len(nums)
	l, r := 0, len
	var pos int
	for l <= r {
		m := (l + r) / 2
		if nums[m] > nums[0] {
			l = m + 1
			pos = m
		} else {
			r = m - 1
		}
	}
	r = pos
	l = (pos + 1) % len
	for l != r {
		m := (l + (r-l+len)%len/2) % len
		if target < nums[m] {
			r = (m - 1 + len) % len
		} else if target > nums[m] {
			l = (m + 1) % len
		} else {
			return m
		}
	}
	if l == r && nums[l] == target {
		return l
	}
	return -1
}
func IsClosed(ch <-chan T) bool {
	select {
	case _, a := <-ch:
		if !a {
			return true
		}
		return false
	default:
	}

	return false
}

func main() {
	c := make(chan T, 2)
	fmt.Println(IsClosed(c)) // false
	c <- 2
	fmt.Println(IsClosed(c))
	close(c)
	fmt.Println(IsClosed(c)) // true
}
