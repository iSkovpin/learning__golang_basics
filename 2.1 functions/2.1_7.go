package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var arr [4]int

	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3])
	var min = arr[0]

	for _, elem := range arr {
		if elem < min {
			min = elem
		}
	}

	return min
}
