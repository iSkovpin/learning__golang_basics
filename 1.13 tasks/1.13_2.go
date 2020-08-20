package main

import "fmt"

func main() {
	var n, sum int

	fmt.Scan(&n)

	sum = n%10 + n%100/10 + n/100

	fmt.Println(sum)
}
