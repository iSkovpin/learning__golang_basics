package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	var h int = d / 30
	var m int = d % 30 * 2

	fmt.Println("It is", h, "hours", m, "minutes.")
}
