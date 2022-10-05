package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := rpnCalculator("5 5 + 6 6 + 1 1 + * +")
	fmt.Println(result)
}

// рекурсивный
// func rpnCalculator(s string) int {

// 	var a, b, result int

// 	if len(s) == 1 {
// 		result, _ = strconv.Atoi(s)
// 		return result
// 	}

// 	expSlice := strings.Fields(s)
// 	for i, v := range expSlice {
// 		if v == "*" {
// 			a, _ = strconv.Atoi(expSlice[i-1])
// 			b, _ = strconv.Atoi(expSlice[i-2])
// 			result = a * b
// 			if len(expSlice) == 3 {
// 				break
// 			}
// 			expSlice[i] = strconv.Itoa(result)
// 			expSlice = append(expSlice[:i-2], expSlice[i:]...)
// 			str := strings.Join(expSlice, " ")
// 			result = rpnCalculator(str)
// 			break
// 		} else if v == "+" {
// 			a, _ = strconv.Atoi(expSlice[i-1])
// 			b, _ = strconv.Atoi(expSlice[i-2])
// 			result = a + b
// 			if len(expSlice) == 3 {
// 				break
// 			}
// 			expSlice[i] = strconv.Itoa(result)
// 			expSlice = append(expSlice[:i-2], expSlice[i:]...)
// 			str := strings.Join(expSlice, " ")
// 			result = rpnCalculator(str)
// 			break
// 		} else {
// 			continue
// 		}
// 	}
// 	return result
// }

func rpnCalculator(s string) int {
	var result int

	expSlice := strings.Fields(s)
	for i := 0; i < len(expSlice); {
		if len(expSlice) == 1 {
			result, _ = strconv.Atoi(expSlice[i])
			break
		}
		switch expSlice[i] {
		case "*":
			a, _ := strconv.Atoi(expSlice[i-1])
			b, _ := strconv.Atoi(expSlice[i-2])
			expSlice[i] = strconv.Itoa(a * b)
			expSlice = append(expSlice[:i-2], expSlice[i:]...)
			i = 0
			continue
		case "+":
			a, _ := strconv.Atoi(expSlice[i-1])
			b, _ := strconv.Atoi(expSlice[i-2])
			expSlice[i] = strconv.Itoa(a + b)
			expSlice = append(expSlice[:i-2], expSlice[i:]...)
			i = 0
			continue
		default:
			i++
			continue
		}
	}
	return result
}
