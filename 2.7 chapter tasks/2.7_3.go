package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(addStars(text))
}

func addStars(text string) string {
	return strings.Join(strings.Split(text, ""), "*")
}
