package main

import "fmt"

func main() {
	s := []int{10, 9, 11, 3, 8, 7, 0}
	fmt.Println(combSort(s))
}

func combSort(s []int) []int {
	n := len(s)
	k := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j+k < n; j++ {
			if s[j] > s[j+k] {
				s[j], s[j+k] = s[j+k], s[j]
			}
		}
		if k > 1 {
			k = int(float64(k) / 1.24733)
		} else {
			break
		}
	}
	return s
}
