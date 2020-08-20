package main

import "fmt"

func main() {
	var sliceLength, positiveCounter int

	fmt.Scan(&sliceLength)

	slice := make([]int, sliceLength)

	for i := 0; i < sliceLength; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < sliceLength; i++ {
		if slice[i] > 0 {
			positiveCounter++
		}
	}

	fmt.Println(positiveCounter)
}
