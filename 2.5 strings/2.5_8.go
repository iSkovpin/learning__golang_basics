package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	runeSlice := []rune(text)

	if isPalindrom(runeSlice) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func isPalindrom(word []rune) bool {
	if len(word) == 0 {
		return false
	}

	for i, j := 0, len(word)-1; i < j; {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}
