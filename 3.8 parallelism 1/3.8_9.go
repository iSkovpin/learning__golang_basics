package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 16)

	go task2(ch, "A")
	go task2(ch, "B")
	go task2(ch, "C")
	time.Sleep(time.Millisecond)
	close(ch)

	printFromChannel(ch)
}

func printFromChannel(ch chan string) {
	for str := range ch {
		fmt.Print(str)
	}
}

func task2(ch chan string, str string) {
	str += " "
	for i := 0; i < 5; i++ {
		ch <- str
	}
}
