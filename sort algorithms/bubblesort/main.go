package main

import "fmt"

func main() {
	s := []int{10, 9, 11, 8, 7}
	fmt.Println(bubblesort(s))
}

func bubblesort(s []int) []int {
	n := len(s)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
