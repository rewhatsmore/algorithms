package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 0}
	fmt.Println(evenOddSort(s))
	fmt.Println(evenOddSort2(s))
}

func evenOddSort(s []int) []int {
	n := len(s)
	sorted := [2]bool{}
	for i := 0; ; i++ {
		sorted[i%2] = true
		for j := i % 2; j < n-1; j += 2 {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				sorted[i%2] = false
			}
		}
		if sorted[0] && sorted[1] {
			break
		}
	}
	return s
}

func evenOddSort2(s []int) []int {
	n := len(s)
	var sorted bool
	for {
		sorted = true
		for i := 0; i < n-1; i += 2 {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				sorted = false
			}
		}
		for i := 1; i < n-1; i += 2 {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
	return s
}
