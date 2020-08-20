package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))
}

func vote(x int, y int, z int) int {
	if x+y+z > 1 {
		return 1
	}
	return 0
}
