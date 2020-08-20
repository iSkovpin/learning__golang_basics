package main

import "fmt"

func main() {
	var a int

	for {
		fmt.Scan(&a)

		if a < 10 {
			continue
		}
		if a > 100 {
			break
		}

		fmt.Println(a)
	}
}
