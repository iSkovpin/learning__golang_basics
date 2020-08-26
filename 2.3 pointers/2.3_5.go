package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	test(&a, &b)
}

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}
