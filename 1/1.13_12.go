package main

import "fmt"

func main() {
	var n int
	var a = 1

	fmt.Scan(&n)

	for a <= n {
		fmt.Print(a, " ")
		a *= 2
	}
}
