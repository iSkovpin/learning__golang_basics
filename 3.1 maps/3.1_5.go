package main

import "fmt"

func main() {
	var number int
	cache := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&number)
		if _, inMap := cache[number]; !inMap {
			cache[number] = work(number)
		}

		fmt.Print(cache[number], " ")
	}
}

func work(x int) int {
	// some logic
	return x
}
