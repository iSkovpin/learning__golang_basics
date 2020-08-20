package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	for i := 1; i <= 3; i++ {
		var first, second int
		fmt.Scan(&first, &second)
		workArray[first], workArray[second] = workArray[second], workArray[first]
	}

	for _, elem := range workArray {
		fmt.Print(elem, " ")
	}
}
