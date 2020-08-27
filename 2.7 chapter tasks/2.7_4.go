package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(findMax(text))
}

func findMax(text string) string {
	var max int32
	for _, elem := range text {
		if elem > max {
			max = elem
		}
	}
	return string(max)
}
