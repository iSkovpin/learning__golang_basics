package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(sumInt(a, b))
}

func sumInt(numbers ...int) (int, int) {
	var sum int

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return len(numbers), sum
}
