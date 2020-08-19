package main

import "fmt"

func main() {
	var n, min, minCounter int

	fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}

	min = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
			minCounter = 1
		} else if numbers[i] == min {
			minCounter++
		}
	}

	fmt.Println(minCounter)
}
