package main

import "fmt"

func main() {
	var arr [5]int
	var max int

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(max)
}
