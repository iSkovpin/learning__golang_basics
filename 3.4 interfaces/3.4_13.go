package main

import "fmt"

func main() {
	var (
		capStr         string
		batteryForTest Battery
	)

	_, err := fmt.Scan(&capStr)
	if err != nil {
		panic("Incorrect input")
	}

	batteryForTest = Battery{}
	batteryForTest.Charge(capStr)
	fmt.Println(batteryForTest)
}

type Battery struct {
	capacity int
}

func (b *Battery) Charge(str string) {
	b.capacity = 0

	for _, char := range str {
		if string(char) == "1" {
			b.capacity++
		}
	}
}

func (b Battery) String() string {
	var result string = "["

	for i := 0; i < 10; i++ {
		if i+b.capacity < 10 {
			result += " "
		} else {
			result += "X"
		}
	}

	result += "]"
	return result
}
