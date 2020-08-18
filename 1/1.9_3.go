package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)

	switch {
	case a < 10:
		b = a
	case a < 100:
		b = a / 10
	case a < 1000:
		b = a / 100
	case a < 10000:
		b = a / 1000
	case a == 10000:
		b = 1
	}

	fmt.Println(b)
}
