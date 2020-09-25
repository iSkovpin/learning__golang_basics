package main

import (
	"fmt"
	"time"
)

func main() {
	arguments := make(chan int, 10)
	done := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(time.Nanosecond)
		timer := time.NewTimer(time.Millisecond)

		for {
			select {
			case <-ticker.C:
				arguments <- 1
			case <-timer.C:
				done <- struct{}{}
				return
			}
		}

	}()

	fmt.Println(<-calculator(arguments, done))
}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)
		var sum int

		for {
			select {
			case x := <-arguments:
				sum += x
			case <-done:
				resultChan <- sum
				return
			}
		}
	}()

	return resultChan
}
