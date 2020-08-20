package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&d)

	a = d % 10
	d = d / 10
	b = d % 10
	d = d / 10
	c = d

	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
