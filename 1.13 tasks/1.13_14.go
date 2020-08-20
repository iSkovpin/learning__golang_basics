package main

import "fmt"

func main() {
	var (
		number int
		fnm1   = 1
		fn     = 2
		n      = 3
		result = -1
	)

	fmt.Scan(&number)

	for fn <= number {
		if number == fn {
			result = n
			break
		}

		fnm1, fn = fn, fnm1+fn
		n++
	}

	fmt.Println(result)
}
