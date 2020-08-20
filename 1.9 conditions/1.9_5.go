package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	if d > 0 {
		fmt.Println("Число положительное")
	} else if d < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
