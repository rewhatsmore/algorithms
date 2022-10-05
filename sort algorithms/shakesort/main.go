package main

import "fmt"

func main() {
	s := []int{10, 9, 11, 3, 8, 7}
	fmt.Println(shakeSort(s))
}

func shakeSort(s []int) []int {
	left := 0
	right := len(s) - 1
	for i := right; i > left; i-- {
		for j := left; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
		for j := i - 1; j > left; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
		left++
	}
	return s
}
