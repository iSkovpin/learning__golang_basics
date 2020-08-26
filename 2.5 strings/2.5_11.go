package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(removeRepeatingChars(text))
}

func removeRepeatingChars(word string) string {
	result := word

	for _, char := range word {
		if strings.Count(word, string(char)) > 1 {
			result = strings.ReplaceAll(result, string(char), "")
		}
	}

	return result
}
