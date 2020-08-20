package main

import "fmt"

func main() {
	var a, b, sum int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)

		if b > -100 && b < 100 && b > 8 && b%8 == 0 {
			sum = sum + b
		}
	}

	fmt.Println(sum)
}
