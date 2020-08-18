package main

import "fmt"

func main() {
	var sliceLength int

	fmt.Scan(&sliceLength)

	slice := make([]int, sliceLength)

	for i := 0; i < sliceLength; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < sliceLength; i = i + 2 {
		fmt.Print(slice[i], " ")
	}
}
