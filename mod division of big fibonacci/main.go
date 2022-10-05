package main

import (
	"fmt"
)

func main() {
	a, b, n, m := 0, 1, 0, 0
	fmt.Scan(&n, &m)
	slice := make([]int, 1, m+1)

	for i := 1; i <= n; i++ {
		a, b = b, (a+b)%m
		slice = append(slice, a)
		if a == 0 && b == 1 {
			n = n % i
			break
		}
	}
	fmt.Println("n = ", n)
	fmt.Println(slice)
	fmt.Println(slice[n])
}

//package main
//
//import (
//"fmt"
//"os"
//"strconv"
//"strings"
//)
//
//func main() {
//	data, _ := os.ReadFile("input.txt")
//
//	input := strings.Fields(string(data))
//	a, _ := strconv.Atoi(input[0])
//	b, _ := strconv.Atoi(input[1])
//
//	output := fmt.Sprintf("%d", a+b)
//	os.WriteFile("output.txt", []byte(output), 0666)
//}
