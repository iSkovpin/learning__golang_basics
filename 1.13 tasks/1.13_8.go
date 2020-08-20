package main

import "fmt"

func main() {
	var n, current, zeros int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&current)
		if current == 0 {
			zeros++
		}
	}

	fmt.Println(zeros)
}
