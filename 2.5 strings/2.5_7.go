package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeSlice := []rune(text)
	runeSlice = runeSlice[:len(runeSlice)-1]

	if isCorrectSentence(runeSlice) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

func isCorrectSentence(runeSlice []rune) bool {
	return len(runeSlice) > 1 && unicode.IsUpper(runeSlice[0]) && runeSlice[len(runeSlice)-1] == '.'
}
