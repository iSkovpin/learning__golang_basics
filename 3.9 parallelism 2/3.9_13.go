package main

import (
	"fmt"
)

func main() {
	firstChan := make(chan int, 1)
	secondChan := make(chan int, 1)
	stopChan := make(chan struct{}, 1)

	firstChan <- 5
	//secondChan <- 5
	//stopChan <- struct{}{}

	fmt.Println(<-calculator(firstChan, secondChan, stopChan))
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		select {
		case x := <-firstChan:
			resultChan <- x * x
		case x := <-secondChan:
			resultChan <- x * 3
		case <-stopChan:
			return
		}
	}()

	return resultChan
}
