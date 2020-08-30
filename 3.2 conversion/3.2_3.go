package main

import "fmt"

func main() {
	var number64 int64
	var number16 uint16

	fmt.Scan(&number64)
	number16 = convert(number64)
	fmt.Printf("%v (%T)", number16, number16)
}

func convert(x int64) uint16 {
	return uint16(x)
}
