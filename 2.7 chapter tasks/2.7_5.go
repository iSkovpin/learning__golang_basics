package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	for _, elem := range sqrtize(number) {
		fmt.Print(elem)
	}
}

func sqrtize(number int) []int {
	var lastDigit int
	var res = make([]int, 0, 20)

	for number > 0 {
		lastDigit = number % 10
		res = append(res, lastDigit*lastDigit)
		number = number / 10
	}

	return reverseInts(res)
}

func reverseInts(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
