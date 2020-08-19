package main

import "fmt"

func main() {
	var n int
	var word string

	fmt.Scan(&n)

	if inRange(n, 11, 19) || inRange(n%10, 5, 9) || n%10 == 0 {
		word = "korov"
	} else if n%10 == 1 {
		word = "korova"
	} else if inRange(n%10, 2, 4) {
		word = "korovy"
	}

	fmt.Printf("%d %s", n, word)
}

func inRange(a int, lft int, rght int) bool {
	for i := lft; i <= rght; i++ {
		if i == a {
			return true
		}
	}
	return false
}
