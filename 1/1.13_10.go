package main

import "fmt"

func main() {
	var a, b, x int

	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			x = i
			break
		}
	}

	if x != 0 {
		fmt.Println(x)
	} else {
		fmt.Println("NO")
	}
}
