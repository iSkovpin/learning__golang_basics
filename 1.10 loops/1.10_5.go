package main

import "fmt"

func main() {
	var a, max, counter int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > max {
			max = a
			counter = 1
		} else if a == max {
			counter++
		}
	}

	fmt.Println(counter)
}
