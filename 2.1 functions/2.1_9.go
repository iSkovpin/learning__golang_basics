package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n < 3 {
		return 1
	}

	var (
		fmm1 = 1
		fm   = 2
		m    = 3
	)

	for {
		if m == n {
			break
		}

		fmm1, fm = fm, fmm1+fm
		m++
	}

	return fm
}
