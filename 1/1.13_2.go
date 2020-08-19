package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d%d%d", n%10, n%100/10, n/100)
}
