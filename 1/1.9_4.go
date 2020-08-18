package main

import "fmt"

func main() {
	var one, two, three, four, five, six, d int
	fmt.Scan(&d)

	six = d % 10
	d = d / 10
	five = d % 10
	d = d / 10
	four = d % 10
	d = d / 10
	three = d % 10
	d = d / 10
	two = d % 10
	d = d / 10
	one = d

	if one+two+three == four+five+six {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
