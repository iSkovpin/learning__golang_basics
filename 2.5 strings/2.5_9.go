package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1, word2 string
	fmt.Scan(&word1, &word2)
	fmt.Println(strings.Index(word1, word2))
}
