package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)
	runeSlice := []rune(password)

	if isCorrectPassword(runeSlice) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isCorrectPassword(password []rune) bool {
	if len(password) < 5 {
		return false
	}

	for _, char := range password {
		if !(unicode.Is(unicode.Latin, char) || unicode.IsDigit(char)) {
			return false
		}
	}

	return true
}
