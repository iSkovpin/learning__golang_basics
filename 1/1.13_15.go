package main

import "fmt"

func main() {
	var srcNum, deleteNum string
	fmt.Scan(&srcNum, &deleteNum)

	for i := 0; i < len(srcNum); i++ {
		if string(srcNum[i]) != deleteNum {
			fmt.Print(string(srcNum[i]))
		}
	}
}
