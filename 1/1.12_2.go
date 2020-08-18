package main

import "fmt"

func main() {
	var sliceLength int

	fmt.Scan(&sliceLength)

	slice := make([]int, 0, sliceLength)

	for i := 0; i < sliceLength; i++ {
		var elem int
		fmt.Scan(&elem)
		slice = append(slice, elem)
	}

	fmt.Println(slice[3])
}
