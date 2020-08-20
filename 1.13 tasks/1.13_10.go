package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for num > 9 {
		num = getSliceSum(sliceNum(num))
	}

	fmt.Println(num)
}

func sliceNum(num int) []int {
	slice := make([]int, 0, 8)

	for num > 0 {
		slice = append(slice, num%10)
		num = num / 10
	}

	return slice
}

func getSliceSum(slice []int) int {
	var sum int

	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	return sum
}
