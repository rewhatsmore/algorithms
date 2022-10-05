package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	for i := 1; ; {
		if a/i != 0 {
			sum += a / i % 10
			i *= 10
			continue
		} else if i == 10 {
			break
		} else {
			a, sum, i = sum, 0, 1
		}
	}
	fmt.Println(sum)
}

//func main() {
//	var a, sum int
//	fmt.Scan(&a)
//	for ; a >= 10; {
//		sum += a % 10
//		a = a / 10
//		if a < 10 {
//			a += sum
//			sum = 0
//		}
//	}
//	fmt.Println(a)
//}

// func main() {
// var a int

//     for fmt.Scan(&a); a/10 > 0; {
//         a = a/10 + a%10
//     }
//     fmt.Println(a)
// }

//func main() {
//	var a int
//	fmt.Scan(&a)
//	fmt.Println((a-1)%9 + 1)
//}
