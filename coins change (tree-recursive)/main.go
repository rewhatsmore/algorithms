package main

import "fmt"

func main() {
	coinTypes := 5
	coinNamination := []int{1, 5, 10, 25, 50}
	count := countChange(100, coinTypes, coinNamination)
	fmt.Println(count)
}

func countChange(amount, coinTypes int, coinNomination []int) (count int) {
	switch {
	case amount == 0:
		count = 1
	case coinTypes == 0 || amount < 0:
		count = 0
	default:
		count = countChange(amount, coinTypes-1, coinNomination[:(len(coinNomination)-1)]) + countChange(amount-coinNomination[len(coinNomination)-1], coinTypes, coinNomination)
	}
	return
}
