package main

import "fmt"

//func main() {
//	var a, b int
//	fmt.Scan(&a, &b)
//	for {
//		if b == 0 {
//			fmt.Println(a)
//			break
//		}
//		if b > a {
//			a, b = b, a
//		}
//		a, b = b, a%b
//	}
//}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b > a {
		a, b = b, a
	}
	fmt.Println(euclid(a, b))
}

func euclid(a, b int) int {
	if b > 0 {
		return euclid(b, a%b)
	}
	return a
}
