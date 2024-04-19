package main

import "fmt"

type T int

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
