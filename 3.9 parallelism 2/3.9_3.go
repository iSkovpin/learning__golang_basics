package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})

	go func(done chan struct{}) {
		work()
		close(done)
	}(done)

	<-done
}

func work() {
	fmt.Println("Work is done")
}
