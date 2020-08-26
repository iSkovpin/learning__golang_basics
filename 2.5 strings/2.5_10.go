package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)
	runeSlice := []rune(word)
	fmt.Println(string(removeOddChars(runeSlice)))
}

func removeOddChars(word []rune) []rune {
	result := make([]rune, 0, len(word)/2)

	for idx, char := range word {
		if idx%2 != 0 {
			result = append(result, char)
		}
	}

	return result
}
