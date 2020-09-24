package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go task(ch, 1)
	go task(ch, 2)
	go task(ch, 3)
	time.Sleep(1000)
	fmt.Println(<-ch)
}

func task(ch chan int, n int) {
	ch <- n + 1
}
