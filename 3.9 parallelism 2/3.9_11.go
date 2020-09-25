package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i // Will not work as you expect
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(len(m))
}
