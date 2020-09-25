package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	var doWork = func(wg *sync.WaitGroup) {
		defer wg.Done()
		work()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(wg)
	}

	wg.Wait()
	fmt.Println("All work is done")
}

func work() {
	fmt.Println("Work is done")
}
