package main

import (
	"fmt"
	"sort"
)

func main() {
	sides := make([]int, 3)
	fmt.Scan(&sides[0], &sides[1], &sides[2])
	sort.Ints(sides)

	if sides[0]+sides[1] > sides[2] {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
